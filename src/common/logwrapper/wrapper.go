package logwrapper

import (
	"fmt"
	"io"
	"os"

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

func CreateLogger(fields map[string]string) (logrus.FieldLogger, error) {
	return entry, nil
}

// GetMyLogger is used by each Model/Function to generate it's own logger instance.
// NOTE: that if there is an error Creating the logger, we panic
func GetMyLogger(requestId string, prefix string) logrus.FieldLogger {
	if len(prefix) == 0 {
		errMsg := "ERROR: Could NOT acquire Logger: prefix value required"
		panic(errMsg)
	}

	stdFlds := map[string]string{
		RequestIdField:      requestId,
		FunctionPrefixField: prefix,
	}
	logger, err := CreateLogger(stdFlds)
	if err != nil {
		msg := "From " + prefix + ": ERROR: Could NOT acquire Logger: " + err.Error()
		fmt.Fprintf(os.Stderr, msg)
		panic(msg)
	}

	return logger
}
