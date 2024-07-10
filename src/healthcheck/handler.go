package healthcheck

import "go.mongodb.org/mongo-driver/mongo"

func NewHandler(config configPkg.Config) Handler {
	return &theHandler{
		config:         config,
		hriHealthcheck: GetCheck,
	}
}

type theHandler struct {
	config         configPkg.Config
	hriHealthcheck func(string, *mongo.Collection) (int, *response.ErrorDetail)
}
