package healthcheck

import (
	configPkg "github.com/VishwasMallikarjuna/go-mongo-appliacation/common/config"
	"github.com/VishwasMallikarjuna/go-mongo-appliacation/common/response"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler interface {
	HriHealthcheck(echo.Context) error
}

func NewHandler(config configPkg.Config) Handler {
	return &theHandler{
		config:      config,
		healthcheck: GetCheck,
	}
}

type theHandler struct {
	config      configPkg.Config
	healthcheck func(string, *mongo.Collection) (int, *response.ErrorDetail)
}
