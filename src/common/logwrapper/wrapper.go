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
