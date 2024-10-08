package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/VishwasMallikarjuna/go-mongo-application/common/config"
	"github.com/VishwasMallikarjuna/go-mongo-application/common/logwrapper"
	"github.com/VishwasMallikarjuna/go-mongo-application/healthcheck"
	mongohandler "github.com/VishwasMallikarjuna/go-mongo-application/mongo"
	mongo "github.com/VishwasMallikarjuna/go-mongo-application/mongoApi"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
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

func logLvlInfoOrLess(logCfg *logwrapper.LogConfig) bool {
	return logCfg.Level == logrus.InfoLevel || logCfg.Level == logrus.DebugLevel ||
		logCfg.Level == logrus.TraceLevel
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
	if logLvlInfoOrLess(logCfg) {
		e.Use(
			middleware.Logger(), // Log every request/response to stdout
		)
	}

	err = mongo.ConnectFromConfig(config)
	if err != nil {
		e.Logger.Fatal(err)
		os.Exit(2)
	}

	// Prepare the server start function
	startFunc := func() {
		err := error(nil)

		err = e.Start(":1323")

		if err != nil {
			e.Logger.Fatal(err)
			os.Exit(2)
		}
	}

	// Configure the endpoint routes

	// Endpoint for ready/liveness probes
	e.GET("/alive", func(c echo.Context) error {
		return c.String(http.StatusOK, "yes")
	})

	// Healthcheck routing
	healthcheckHandler := healthcheck.NewHandler(config)
	e.GET("/healthcheck", healthcheckHandler.Healthcheck)

	mongoHandler := mongohandler.NewHandler(config)
	e.POST("/users", mongoHandler.Create)
	e.GET("/users", mongoHandler.Get)
	e.PUT("/users", mongoHandler.Update)
	e.DELETE("/users", mongoHandler.Delete)
	return 0, startFunc, nil
}
