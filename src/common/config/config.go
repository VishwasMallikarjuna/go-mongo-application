package config

import "flag"

type Config struct {
	ConfigPath   string
	LogLevel     string
	MongoDBUri   string
	MongoDBName  string
	MongoColName string
}

func GetConfig(configPath string, commandLineFlags []string) (Config, error) {
	fs := flag.NewFlagSet("goMongoConfig", flag.ContinueOnError)

	config := Config{}
	fs.StringVar(&config.ConfigPath, "config-path", configPath, "(Optional) Path of an alternate config file")

	return config, nil
}
