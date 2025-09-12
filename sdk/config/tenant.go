package config

type HostTenant struct {
	Enabled bool      `yaml:"enabled"`
	Id      uint64    `yaml:"id"`
	Key     string    `yaml:"key"`
	Db      *Database `yaml:"db"`
}

var (
	HostTenantsConfig = make(map[string]*HostTenant)
)
