package delivery

import (
	"encoding/json"
	"net/http"
)

type CustomJSONUtil struct{}

type CustomJSONUtilInterface interface {
	CustomJSONRes(w http.ResponseWriter, key string, value string, httpStatus int, data interface{})
}

func (cj *CustomJSONUtil) CustomJSONRes(w http.ResponseWriter, key string, value string, httpStatus int, data interface{}) {
	w.Header().Set(key, value)
	w.WriteHeader(httpStatus)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
