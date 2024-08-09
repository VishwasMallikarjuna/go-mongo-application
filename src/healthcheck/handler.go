package healthcheck

import (
	"net/http"

	configPkg "github.com/VishwasMallikarjuna/go-mongo-application/common/config"
	"github.com/VishwasMallikarjuna/go-mongo-application/common/logwrapper"
	"github.com/VishwasMallikarjuna/go-mongo-application/common/response"
	"github.com/VishwasMallikarjuna/go-mongo-application/mongoApi"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
)

type theHandler struct {
	config      configPkg.Config
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

	requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	prefix := "healthcheck/handler"
	var logger = logwrapper.GetMyLogger(requestId, prefix)
	logger.Debug("Start Healthcheck Handler")

	code, errorDetail := h.healthcheck(requestId, mongoApi.GetMongoCollection(h.config.MongoColName))
	if errorDetail != nil {
		return c.JSON(http.StatusInternalServerError, errorDetail)
	}

	return c.NoContent(code)
}
