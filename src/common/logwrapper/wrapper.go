package logwrapper

import (
	"fmt"
	"io"

	"github.com/sirupsen/logrus"
)

const (
	RequestIdField      = "requestId"
	FunctionPrefixField = "functionPrefix"
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

func ParseLevelFromStr(lvlStr string) (logrus.Level, error) {
	level, err := logrus.ParseLevel(lvlStr)
	if err != nil {
		wrappedErr := fmt.Errorf("error parsing log Level - %w", err)
		return logrus.PanicLevel, wrappedErr ///in logrus, Default Level => Panic
	}
	return level, nil
}
