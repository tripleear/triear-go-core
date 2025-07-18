package mycasbin

import (
	"context"
	"regexp"
	"strings"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/log"
	"github.com/casbin/casbin/v2/model"
	redisWatcher "github.com/nyl1001/redis-watcher/v2"
	redis "github.com/redis/go-redis/v9"
	"github.com/tripleear/triear-go-core/logger"
	"github.com/tripleear/triear-go-core/sdk"
	"github.com/tripleear/triear-go-core/sdk/config"
	"gorm.io/gorm"

	gormAdapter "github.com/nyl1001/gorm-adapter/v3"
)

// Initialize the model from a string.
var textOld = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && (keyMatch2(r.obj, p.obj) || keyMatch(r.obj, p.obj)) && (r.act == p.act || p.act == "*")
`

var text = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && customMatch(r.obj, p.obj) && (r.act == p.act || p.act == "*")
`

var (
	enforcer   *casbin.SyncedEnforcer
	once       sync.Once
	digitRegex = regexp.MustCompile(`^\d+$`)
)

// 自定义匹配函数：支持 /:id 但要求实际值为纯数字
func customMatch(requestPath, policyPath string) bool {
	if requestPath == policyPath {
		return true
	}

	reqSegs := strings.Split(strings.Trim(requestPath, "/"), "/")
	polSegs := strings.Split(strings.Trim(policyPath, "/"), "/")

	if len(reqSegs) != len(polSegs) {
		return false
	}

	for i := 0; i < len(reqSegs); i++ {
		if strings.HasPrefix(polSegs[i], ":") {
			if !digitRegex.MatchString(reqSegs[i]) {
				return false
			}
		} else if reqSegs[i] != polSegs[i] {
			return false
		}
	}
	return true
}

func CustomMatchFunc(args ...interface{}) (interface{}, error) {
	reqPath := args[0].(string)
	polPath := args[1].(string)
	return customMatch(reqPath, polPath), nil
}

func Setup(db *gorm.DB, _ string) *casbin.SyncedEnforcer {
	once.Do(func() {
		Apter, err := gormAdapter.NewAdapterByDBUseTableName(db, "", "casbin_rule")
		if err != nil && err.Error() != "invalid DDL" {
			panic(err)
		}

		m, err := model.NewModelFromString(text)
		if err != nil {
			panic(err)
		}
		enforcer, err = casbin.NewSyncedEnforcer(m, Apter)
		if err != nil {
			panic(err)
		}
		enforcer.AddFunction("customMatch", CustomMatchFunc)
		err = enforcer.LoadPolicy()
		if err != nil {
			panic(err)
		}
		// set redis watcher if redis config is not nil
		if config.CacheConfig.Redis != nil {
			w, err := redisWatcher.NewWatcher(config.CacheConfig.Redis.Addr, redisWatcher.WatcherOptions{
				Options: redis.Options{
					Network:  "tcp",
					Password: config.CacheConfig.Redis.Password,
				},
				Channel:    "/casbin",
				IgnoreSelf: false,
			})
			if err != nil {
				panic(err)
			}

			err = w.SetUpdateCallback(updateCallback)
			if err != nil {
				panic(err)
			}
			err = enforcer.SetWatcher(w)
			if err != nil {
				panic(err)
			}
		}

		log.SetLogger(&Logger{})
		enforcer.EnableLog(true)
	})

	return enforcer
}

func updateCallback(msg string) {
	l := logger.NewHelper(sdk.Runtime.GetLogger())
	ctx := context.Background()
	l.Infof(ctx, "casbin updateCallback msg: %v", msg)
	err := enforcer.LoadPolicy()
	if err != nil {
		l.Errorf(ctx, err, "casbin LoadPolicy err: %v")
	}
}
