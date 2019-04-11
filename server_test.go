package pocket

import "testing"
import "github.com/stretchr/testify/require"

func TestServer(t *testing.T) {

	handler := NewRequestHandler()

	handler.Post("/posttest", func(responder Responder) error {
		return nil
	})

	handler.Get("/gettest", func(responder Responder) error {
		return nil
	})

	server, err := NewServer()
	require.NoError(t, err)

	require.NoError(t, server.SetHandler(handler))

	require.NoError(t, server.Stop())

}
