package logwrapper

import (
	"io"

	"github.com/sirupsen/logrus"
)

type LogConfig struct {
	Level    logrus.Level
	Location io.Writer
}

var globalConfig LogConfig

func Initialize(lvlStr string, out io.Writer) (*LogConfig, error) {
	parsedLvl, err := ParseLevelFromStr(lvlStr)
	if err != nil {
		return nil, err
	}

	globalConfig = LogConfig{
		Level:    parsedLvl,
		Location: out,
	}
	return &globalConfig, nil
}
