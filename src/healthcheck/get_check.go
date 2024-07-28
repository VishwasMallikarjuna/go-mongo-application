package healthcheck

import (
	"github.com/VishwasMallikarjuna/go-mongo-appliacation/common/logwrapper"
	"github.com/VishwasMallikarjuna/go-mongo-appliacation/common/response"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCheck(requestId string, client *mongo.Collection) (int, *response.ErrorDetail) {
	prefix := "healthcheck/getCheck"
	var logger = logwrapper.GetMyLogger(requestId, prefix)
	logger.Infof("Prepare HealthCheck - mongDb (No Input Params)")
}
