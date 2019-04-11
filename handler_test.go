package pocket

import "testing"

func TestHandler(t *testing.T) {

	handler := NewRequestHandler()

	handler.Post("/posttest", func(responder Responder) error {
		return nil
	})

	handler.Get("/gettest", func(responder Responder) error {
		return nil
	})

}
