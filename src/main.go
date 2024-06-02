package main

import (
	"fmt"
	"os"

	"github.com/VishwasMallikarjuna/go-mongo-appliacation/common/config"
	"github.com/labstack/echo"
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

	return 0, startFunc, nil
}
