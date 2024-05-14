package config

import "flag"

type Config struct {
	ConfigPath         string
	LogLevel           string
	MongoDBUri         string
	MongoDBName        string
	MongoColName       string
}

func GetConfig(configPath string, commandLineFlags []string) (Config, error) {
	fs := flag.NewFlagSet("goMongoConfig", flag.ContinueOnError)

	config := Config{}

}
