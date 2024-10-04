package config

import (
	"github.com/joomcode/errorx"
	"github.com/spf13/viper"
)

var (
	Environment string
)

var (
	ConfigErr    = errorx.NewNamespace("config")
	FileReadErr  = ConfigErr.NewType("file_read_err")
	UnmarshalErr = ConfigErr.NewType("unmarshal_err")
)

func Load(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, FileReadErr.Wrap(err, "failed to read config file")
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, UnmarshalErr.Wrap(err, "failed to unmarshal config")
	}

	return &config, nil
}
