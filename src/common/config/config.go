package config

import (
	"errors"
	"flag"
	"strings"

	"github.com/peterbourgon/ff/v3"
	"github.com/peterbourgon/ff/v3/ffyaml"
)

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

func ValidateConfig(config Config) error {

	if len(config.ConfigPath) == 0 {
		return errors.New("no config file supplied")
	}

	errorBuilder := strings.Builder{}
	errorHeader := "Configuration errors:"
	errorBuilder.WriteString(errorHeader)

	if config.NewRelicEnabled && config.NewRelicAppName == "" {
		errorBuilder.WriteString("\n\tNew Relic monitoring enabled, but the New Relic app name was not specified")
	}

	if config.NewRelicEnabled && config.NewRelicLicenseKey == "" {
		errorBuilder.WriteString("\n\tNew Relic monitoring enabled, but the New Relic license key was not specified")
	}
	return nil
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

	err := ff.Parse(fs, commandLineFlags,
		ff.WithIgnoreUndefined(true),
		ff.WithConfigFileFlag("config-path"),
		ff.WithConfigFileParser(ffyaml.Parser),
		ff.WithAllowMissingConfigFile(false),
		ff.WithEnvVarNoPrefix(),
	)
	if err != nil {
		// If the issue isn't related to CL args, also print the usage guide
		// (it will get printed automatically otherwise).
		if !strings.Contains(err.Error(), "error parsing commandline args: invalid ") {
			fs.Usage()
		}
		return config, err
	}

	err = ValidateConfig(config)
	if err != nil {
		fs.Usage()
		return Config{}, err
	}

	return config, nil
}
