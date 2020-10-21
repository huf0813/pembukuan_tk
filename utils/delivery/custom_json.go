package delivery

import (
	"encoding/json"
	"github.com/huf0813/pembukuan_tk/model"
	"net/http"
)

type CustomJSONUtil struct {
	Payload model.Payload
}

type CustomJSONUtilInterface interface {
	CustomJSONRes(w http.ResponseWriter, key string, value string, httpStatus int, payload model.Payload)
}

func (cj *CustomJSONUtil) CustomJSONRes(w http.ResponseWriter, key string, value string, httpStatus int, status, message string, data interface{}) {
	w.Header().Set(key, value)
	w.WriteHeader(httpStatus)
	if err := json.NewEncoder(w).Encode(cj.Payload.NewPayload(status, message, data)); err != nil {
		panic(err)
	}
}
