package pocket

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ResponseHandler func(Responder) error

type Responder struct {
	method  string
	path    string
	writer  http.ResponseWriter
	request *http.Request
}

func NewResponder(method string, path string) *Responder {
	return &Responder{method: method, path: path}
}

func (responder *Responder) GetMethod() string {
	return responder.method
}

func (responder *Responder) GetPath() string {
	return responder.path
}

func (responder *Responder) Prepare(writer http.ResponseWriter, request *http.Request) (Responder, error) {

	copy := *responder

	if writer == nil {
		return copy, errors.New("writer is nil")
	}

	if request == nil {
		return copy, errors.New("request is nil")
	}

	copy.writer = writer
	copy.request = request

	return copy, nil

}

func (responder *Responder) JSON(i interface{}) error {

	encoded, err := json.Marshal(true)
	if err != nil {
		return err
	}

	responder.writer.Header().Set("Content-Type", "application/json")

	_, err = responder.writer.Write(encoded)

	return err
}
