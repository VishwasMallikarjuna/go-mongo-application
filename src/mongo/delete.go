package mongo

import "net/http"

func Delete(requestId string, deleteID string) (int, interface{}) {
	return http.StatusOK, nil
}
