package config

type HostTenants struct {
	Tenants []*Tenant `yaml:"tenants"`
	Db      *Database `yaml:"db"`
}

type Tenant struct {
	Id  uint64 `yaml:"id"`
	Key string `yaml:"key"`
}

var (
	HostTenantsConfig = make(map[string]*HostTenants)
)
