package config

type ApplicationConfig struct {
	envValues *envConfig
}

func (cfg *ApplicationConfig) Version() string {
	return cfg.envValues.Version
}

func (cfg *ApplicationConfig) ServerPort() int {
	return cfg.envValues.ServerPort
}

func (cfg *ApplicationConfig) BaseURL() string {
	return cfg.envValues.BaseUrl
}

func NewApplicationConfig() *ApplicationConfig {
	envValues := newEnvironmentConfig()

	return &ApplicationConfig{
		envValues: envValues,
	}
}
