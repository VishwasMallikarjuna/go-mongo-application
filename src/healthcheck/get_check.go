package healthcheck

import (
	"net/http"

	"github.com/VishwasMallikarjuna/go-mongo-appliacation/common/logwrapper"
	"github.com/VishwasMallikarjuna/go-mongo-appliacation/common/response"
	"github.com/VishwasMallikarjuna/go-mongo-appliacation/mongoApi"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCheck(requestId string, client *mongo.Collection) (int, *response.ErrorDetail) {
	prefix := "healthcheck/getCheck"
	var logger = logwrapper.GetMyLogger(requestId, prefix)
	logger.Infof("Prepare HealthCheck - mongDb (No Input Params)")

	health_status, _, err := mongoApi.DatabaseHealthCheck(client)

	if err != nil {
		return http.StatusServiceUnavailable, mongoApi.LogAndBuildErrorDetail(requestId, http.StatusServiceUnavailable, logger, "Could not perform Cosmos health check")
	}

	var mErrMsg = ""
	if health_status != "1" {
		mErrMsg = ServiceUnavailableMsg
		logger.Errorln(mErrMsg)
		return http.StatusServiceUnavailable, response.NewErrorDetail(requestId, mErrMsg)

	}

	return http.StatusOK, nil
}
