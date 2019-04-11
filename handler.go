package pocket

import (
	"net/http"
)

type RequestHandler struct {
	responders map[*Responder]ResponseHandler
}

func NewRequestHandler() *RequestHandler {
	return &RequestHandler{make(map[*Responder]ResponseHandler)}
}

func (handler *RequestHandler) HasResponder(method string, path string) bool {

	for responder := range handler.responders {

		if method == responder.GetMethod() && path == responder.GetPath() {
			return true
		}

	}

	return false

}

func (handler *RequestHandler) Handle(method string, path string, handle ResponseHandler) {

	var responder *Responder

	for r := range handler.responders {

		if method == r.GetMethod() && path == r.GetPath() {
			responder = r
			break
		}

	}

	if responder == nil {
		responder = NewResponder(method, path)
	}

	handler.responders[responder] = handle

}

func (handler *RequestHandler) Post(path string, handle func(Responder) error) {
	handler.Handle(http.MethodPost, path, handle)
}

func (handler *RequestHandler) Get(path string, handle func(Responder) error) {
	handler.Handle(http.MethodGet, path, handle)
}

func (handler *RequestHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	for responder, handle := range handler.responders {

		if request.Method == responder.GetMethod() && request.URL.Path == responder.GetPath() {

			copy, err := responder.Prepare(writer, request)
			if err != nil {
				panic(err)
			}

			handle(copy)

		}

	}

}
