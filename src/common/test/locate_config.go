package test

import (
	"flag"
	"fmt"
	"testing"

	"github.com/peterbourgon/ff"
	"github.com/peterbourgon/ff/ffyaml"
)

func FindConfigPath(t *testing.T) string {
	possiblePaths := []string{
		"./src/common/config/goodConfig.yml",
		"../config/goodConfig.yml",
		"./goodConfig.yml",
		"common/config/goodConfig.yml",
	}
	fs := flag.NewFlagSet("servetest", flag.ContinueOnError)
	var args []string
	opts := []ff.Option{ff.WithConfigFileParser(ffyaml.Parser)}
	for _, path := range possiblePaths {
		opts = append(opts, ff.WithConfigFile(path))
		err := ff.Parse(fs, args, opts...)
		if err.Error() != fmt.Sprintf("open %s: no such file or directory", path) {
			return path
		}
	}
	t.Fatalf("COULD NOT FIND A CONFIG FILE!")
	return ""
}
