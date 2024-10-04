package config

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

type HttpClientConfig struct {
	TimeoutSec int    `mapstructure:"timeout_sec"`
	Retries    int    `mapstructure:"retries"`
	ProxyURL   string `mapstructure:"proxy_url"`
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"db_name"`
	SSLMode  string `mapstructure:"ssl_mode"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type Config struct {
	Environment      string            `mapstructure:"environment"`
	ServerConfig     *ServerConfig     `mapstructure:"server"`
	HttpClientConfig *HttpClientConfig `mapstructure:"http_client"`
	DBConfig         *DBConfig         `mapstructure:"db"`
	RedisConfig      *RedisConfig      `mapstructure:"redis"`
}
