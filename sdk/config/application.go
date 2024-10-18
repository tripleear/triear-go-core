package config

type Application struct {
	ReadTimeout       int
	WriterTimeout     int
	Host              string
	Port              int64
	Name              string
	Mode              string
	DemoMsg           string
	EnableDP          bool
	EnableSwaggerPage bool `json:"enableSwaggerPage" yaml:"enableSwaggerPage"`
	EnableFormPage    bool `json:"enableFormPage" yaml:"enableFormPage"`
}

var ApplicationConfig = new(Application)
