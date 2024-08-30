package mongo

import "net/http"

func Create(
	requestId string,
	createId string) (int, interface{}) {
	return http.StatusCreated, nil
}
