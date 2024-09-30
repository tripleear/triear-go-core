package config

type S3Config struct {
	Cos []CosConfig `yaml:"cos" json:"cos"`
}

type CosConfig struct {
	Name     string `yaml:"name" json:"name"`
	Endpoint string `yaml:"endpoint" json:"endpoint"`
	Ak       string `yaml:"ak" json:"ak"`
	Sk       string `yaml:"sk" json:"sk"`
	Bucket   string `yaml:"bucket" json:"bucket"`
}

var (
	S3 = new(S3Config)
)
