package config

import "flag"

type Config struct {
}

func GetConfig(configPath string, commandLineFlags []string) (Config, error) {
	fs := flag.NewFlagSet("goMongoConfig", flag.ContinueOnError)

	config := Config{}

}
