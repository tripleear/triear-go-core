package config

type Database struct {
	Driver          string
	Source          string
	ConnMaxIdleTime int
	ConnMaxLifeTime int
	MaxIdleConns    int
	MaxOpenConns    int
	CachePrefix     string
	EsPrefix        string
	Registers       []DBResolverConfig
}

type DBResolverConfig struct {
	Sources  []string
	Replicas []string
	Policy   string
	Tables   []string
}

var (
	PlatformDatabaseConfig = new(Database)
	DatabaseConfig         = new(Database)
	DatabasesConfig        = make(map[string]*Database)
)
