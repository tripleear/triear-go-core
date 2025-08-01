package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tripleear/triear-go-core/sdk/pkg"
	jwt "github.com/tripleear/triear-go-core/sdk/pkg/jwtauth"
)

func ExtractClaims(c *gin.Context) jwt.MapClaims {
	claims, exists := c.Get(jwt.JwtPayloadKey)
	if !exists {
		return make(jwt.MapClaims)
	}

	return claims.(jwt.MapClaims)
}

func Get(c *gin.Context, key string) interface{} {
	data := ExtractClaims(c)
	if data[key] != nil {
		return data[key]
	}

	fmt.Println(pkg.GetCurrentTimeStr() + " [WARING] " + c.Request.Method + " " + c.Request.URL.Path + " Get 缺少 " + key)

	return nil
}

// GetUserId 获取一个int的userId
func GetUserId(c *gin.Context) int {
	data := ExtractClaims(c)
	identity, err := data.Identity()
	if err != nil {
		fmt.Println(pkg.GetCurrentTimeStr() + " [WARING] " + c.Request.Method + " " + c.Request.URL.Path + " GetUserId 缺少 identity error: " + err.Error())
		return 0
	}

	return int(identity)
}

// GetUserIdInt64 获得int64的userId
func GetUserIdInt64(c *gin.Context) int64 {
	data := ExtractClaims(c)
	identity, err := data.Identity()
	if err != nil {
		fmt.Println(pkg.GetCurrentTimeStr() + " [WARING] " + c.Request.Method + " " + c.Request.URL.Path + " GetUserId 缺少 identity error: " + err.Error())
		return 0
	}

	return identity
}

func GetUserIdStr(c *gin.Context) string {
	data := ExtractClaims(c)

	return data.String("identity")
}

func GetUserName(c *gin.Context) string {
	return ExtractClaims(c).String("nice")
}

func GetRoleKeys(c *gin.Context) []string {
	return ExtractClaims(c).StringArray(jwt.RoleKeys)
}

func GetRoleIds(c *gin.Context) []int {
	roleIds, err := ExtractClaims(c).IntArray(jwt.RoleIdKeys)
	if err != nil {
		fmt.Println(pkg.GetCurrentTimeStr() + " [WARING] " + c.Request.Method + " " + c.Request.URL.Path + " GetRoleId 缺少 roleids error: " + err.Error())
		return nil
	}
	return roleIds
}

func GetDeptId(c *gin.Context) int {
	deptId, err := ExtractClaims(c).Int("deptid")
	if err != nil {
		fmt.Println(pkg.GetCurrentTimeStr() + " [WARING] " + c.Request.Method + " " + c.Request.URL.Path + " GetDeptId 缺少 deptid error: " + err.Error())
		return 0
	}

	return deptId
}

func GetDeptName(c *gin.Context) string {
	return ExtractClaims(c).String("deptkey")
}
