package config

type HostTenant struct {
	Enabled bool      `yaml:"enabled"`
	Id      uint64    `yaml:"id"`
	Key     string    `yaml:"key"`
	Code    string    `yaml:"code"`
	Db      *Database `yaml:"db"`
}

var (
	HostTenantsConfig = make(map[string]*HostTenant)
)
