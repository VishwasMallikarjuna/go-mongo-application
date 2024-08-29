package mongo

import "net/http"

func Create(
	requestId string,
	tenantId string) (int, interface{}) {
	return http.StatusCreated, nil
}
