package healthcheck

import (
	configPkg "github.com/VishwasMallikarjuna/go-mongo-appliacation/common/config"
	"github.com/VishwasMallikarjuna/go-mongo-appliacation/common/response"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
)

type theHandler struct {
	config         configPkg.Config
	healthcheck func(string, *mongo.Collection) (int, *response.ErrorDetail)
}

type Handler interface {
	Healthcheck(echo.Context) error
}

func NewHandler(config configPkg.Config) Handler {
	return &theHandler{
		config:      config,
		healthcheck: GetCheck,
	}
}

func (h *theHandler) Healthcheck(c echo.Context) error {
	
	return c.NoContent(code)
}
