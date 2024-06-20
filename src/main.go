package main

import (
	"fmt"
	"os"

	"github.com/VishwasMallikarjuna/go-mongo-appliacation/common/config"
	"github.com/VishwasMallikarjuna/go-mongo-appliacation/common/logwrapper"
	mongo "github.com/VishwasMallikarjuna/go-mongo-appliacation/mongoApi"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	retCode, startServer, _ := configureMgmtServer(e, os.Args[1:])
	if retCode != 0 {
		os.Exit(retCode)
	}

	startServer()
	os.Exit(0)
}

func configureMgmtServer(e *echo.Echo, args []string) (int, func(), error) {
	configPath := "./config.yml"

	config, err := config.GetConfig(configPath, args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR CREATING CONFIG: %v\n", err)
		return 1, nil, err
	}

	//Initialize Logging - NOTE: we are hard-coding the log output location to Stdout
	logCfg, err := logwrapper.Initialize(config.LogLevel, os.Stdout)
	if err != nil {
		msg := "ERROR: Could NOT initialize Logger: %w"
		fmt.Fprintf(os.Stderr, fmt.Errorf(msg, err).Error()+"\n")
		return 3, nil, err //special return code for logging problems
	}
	stdFlds := map[string]string{
		logwrapper.RequestIdField:      "",
		logwrapper.FunctionPrefixField: "SERVE",
	}

	//Log the Server startup
	logger, err := logwrapper.CreateLogger(stdFlds)
	if err != nil {
		msg := "ERROR: Could NOT acquire Logger: " + err.Error()
		fmt.Fprintf(os.Stderr, msg)
		return 3, nil, err //special return code for logging problems
	}
	logger.Infof("serve Startup with minimum Log Level: %s", config.LogLevel)

	e.Use(
		middleware.RequestID(), // Generate a request id on the HTTP response headers
	)

	err = mongo.ConnectFromConfig(config)
	if err != nil {
		e.Logger.Fatal(err)
		os.Exit(2)
	}

	return 0, startFunc, nil
}
