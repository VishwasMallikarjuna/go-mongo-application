package healthcheck

import (
	configPkg "github.com/VishwasMallikarjuna/go-mongo-appliacation/common/config"
	"go.mongodb.org/mongo-driver/mongo"
)

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
