package main

import (
	"go-boiler-plate/config"

	"github.com/projectdiscovery/goflags"
)

// Options contains the command line options
type Options struct {
	// ConfigFilePath is the path to the config file
	ConfigFilePath string
	// Port is the port to listen on
	Port int
}

func parseFlags() *config.Config {
	var options Options
	flagSet := goflags.NewFlagSet()

	// Add the flags to the flagset
	flagSet.CreateGroup("Config", "Config",
		flagSet.StringVarP(&options.ConfigFilePath, "config-file", "cf", "config.yaml", "Config file path"),
		flagSet.DynamicVar(&options.Port, "port", 8080, "Server port to listen on"),
	)

	if err := flagSet.Parse(); err != nil {
		panic(err)
	}
	var cfg *config.Config
	var err error
	cfg, err = config.Load(options.ConfigFilePath)
	if err != nil {
		panic(err)
	}

	// Override the port in config if set in flags
	if options.Port != 0 {
		cfg.ServerConfig.Port = options.Port
	}

	return cfg
}
