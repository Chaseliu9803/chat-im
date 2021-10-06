package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type Msg struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}

func Fail(w http.ResponseWriter, msg string) {
	Resp(w, -1, nil, msg)
}

func Success(w http.ResponseWriter, data interface{}, msg string) {
	Resp(w, 0, data, msg)
}

func Resp(w http.ResponseWriter, code int, data interface{}, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	respMsg := Msg{
		Code: code,
		Message: msg,
		Data: data,
	}

	respJs, err := json.Marshal(respMsg)
	if err != nil {
		log.Println(err.Error())
	}
	w.Write(respJs)
}
