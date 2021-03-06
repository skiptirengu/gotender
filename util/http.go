package util

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type HttpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewHttpError(w http.ResponseWriter, code int) (HttpError) {
	message := http.StatusText(code)
	return NewHttpErrorWithMessage(w, code, message)
}

func NewHttpErrorWithMessage(w http.ResponseWriter, code int, message string) (HttpError) {
	err := HttpError{Code: code}
	err.Message = message

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)

	fmt.Fprint(w, err.ToJSON())
	return err
}

func (h HttpError) ToJSON() (string) {
	res, _ := json.Marshal(&h)
	return string(res)
}
