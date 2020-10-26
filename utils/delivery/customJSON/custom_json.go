package customJSON

import (
	"encoding/json"
	"github.com/huf0813/pembukuan_tk/model"
	"net/http"
)

type JSONCustom struct {
	Payload model.Payload
}

type JSONCustomInterspace interface {
	CustomJSONRes(w http.ResponseWriter, key string, value string, httpStatus int, status, message string, data interface{})
}

func (cj *JSONCustom) CustomJSONRes(w http.ResponseWriter, key string, value string, httpStatus int, status, message string, data interface{}) {
	w.Header().Set(key, value)
	w.WriteHeader(httpStatus)
	if err := json.NewEncoder(w).Encode(cj.Payload.NewPayload(status, message, data)); err != nil {
		panic(err)
	}
}
