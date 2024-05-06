package main

import (
	"os"

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
