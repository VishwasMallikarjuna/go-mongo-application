package healthcheck

import "go.mongodb.org/mongo-driver/mongo"

func GetCheck(requestId string, client *mongo.Collection) (int, *response.ErrorDetail) {

}
