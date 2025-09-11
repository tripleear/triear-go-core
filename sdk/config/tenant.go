package config

type HostTenants struct {
	Tenants map[string]*Tenant `yaml:"tenants"`
	Db      *Database          `yaml:"db"`
}

type Tenant struct {
	Id  uint64 `yaml:"id"`
	Key string `yaml:"key"`
}

var (
	HostTenantsConfig = new(HostTenants)
)
