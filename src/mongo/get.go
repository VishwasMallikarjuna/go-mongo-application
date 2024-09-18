package mongo

import "net/http"

func Get(requestId string, createId string) (int, interface{}) {
	return http.StatusOK, nil
}
