package config

type HostTenant struct {
	Id  uint64    `yaml:"id"`
	Key string    `yaml:"key"`
	Db  *Database `yaml:"db"`
}

var (
	HostTenantConfig = make(map[string]*HostTenant)
)
