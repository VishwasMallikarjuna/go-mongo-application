package healthcheck

import (
	"github.com/VishwasMallikarjuna/go-mongo-appliacation/common/response"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCheck(requestId string, client *mongo.Collection) (int, *response.ErrorDetail) {

}
