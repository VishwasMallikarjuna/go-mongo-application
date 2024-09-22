package mongo

import (
	"net/http"

	"github.com/VishwasMallikarjuna/go-mongo-application/common/config"
	"github.com/labstack/echo"
)

type Handler interface {
	Create(echo.Context) error
	Get(echo.Context) error
	Update(echo.Context) error
}

type theHandler struct {
	config config.Config
	create func(string, string) (int, interface{})
	get    func(string, string) (int, interface{})
}

func NewHandler(config config.Config) Handler {
	return &theHandler{
		config: config,
		create: Create,
		get:    Get,
	}
}

func (h *theHandler) Create(c echo.Context) error {
	return c.JSON(http.StatusCreated, nil)
}

func (h *theHandler) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}
