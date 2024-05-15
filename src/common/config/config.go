package config

import "flag"

type Config struct {
	ConfigPath         string
	LogLevel           string
	NewRelicEnabled    bool
	NewRelicAppName    string
	NewRelicLicenseKey string
	MongoDBUri         string
	MongoDBName        string
	MongoColName       string
}

func GetConfig(configPath string, commandLineFlags []string) (Config, error) {
	fs := flag.NewFlagSet("goMongoConfig", flag.ContinueOnError)

	config := Config{}
	fs.StringVar(&config.ConfigPath, "config-path", configPath, "(Optional) Path of an alternate config file")
	fs.StringVar(&config.LogLevel, "log-level", "info", "(Optional) Minimum Log Level for logging output. Available levels are: Trace, Debug, Info, Warning, Error, Fatal and Panic.")
	fs.BoolVar(&config.NewRelicEnabled, "new-relic-enabled", false, "(Optional) True to enable New Relic monitoring, false otherwise")
	fs.StringVar(&config.NewRelicAppName, "new-relic-app-name", "", "(Optional) Application name to aggregate data under in New Relic")
	fs.StringVar(&config.NewRelicLicenseKey, "new-relic-license-key", "", "(Optional) New Relic license key")
	fs.StringVar(&config.MongoDBUri, "mongoDb-uri", "", "(Optional) Azure cosmosDB Mongo API uri")
	fs.StringVar(&config.MongoDBName, "mongoDb-name", "", "(Optional) Azure cosmosDB Mongo Database name")
	fs.StringVar(&config.MongoColName, "mongoCol-name", "", "(Optional) Azure cosmosDB Mongo Database Collection name")
	return config, nil
}
