package pocket

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
)

type Server struct {
	server *http.Server
}

func IsPortOpen(port int) bool {

	if ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port)); err == nil {

		if err := ln.Close(); err != nil {
			panic(err)
		}

		return true

	}

	return false

}

func NewServer() (*Server, error) {

	if !IsPortOpen(8080) {
		log.Printf("Port %d is not available!", 8080)
		return nil, fmt.Errorf("port %d is not available", 8080)
	}

	return &Server{&http.Server{Addr: fmt.Sprintf(":%d", 8080), Handler: NewRequestHandler()}}, nil
}

func NewServerOnPort(port int) (*Server, error) {

	if !IsPortOpen(port) {
		log.Printf("Port %d is not available!", port)
		return nil, fmt.Errorf("port %d is not available", port)
	}

	return &Server{&http.Server{}}, nil
}

func (server *Server) SetHandler(handler *RequestHandler) error {

	if handler == nil {
		return errors.New("handler is nil")
	}

	server.server.Handler = handler

	return nil

}

func (server *Server) Stop() error {
	return server.server.Shutdown(context.Background())
}
