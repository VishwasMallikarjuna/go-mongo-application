package config

import "flag"

func GetConfig(configPath string, commandLineFlags []string) (Config, error) {
	fs := flag.NewFlagSet("goMongoConfig", flag.ContinueOnError)

}
