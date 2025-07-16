package runtime

import "gorm.io/gorm"

type Context struct {
	Orm  *gorm.DB
	Host string
}

type CacheContext struct {
	Context
	CachePrefix string
	EsPrefix    string
}

type EsContext struct {
	Context
	CachePrefix string
	EsPrefix    string
}
