package mongo

import "net/http"

func Update(requestId string, createId string) (int, interface{}) {
	return http.StatusCreated, nil
}
