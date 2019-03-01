package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoute(t *testing.T) {
	addr := "unix:///tmp/test.sock"
	gsv := startGrpc(addr)
	defer gsv.Stop()

	t.Run("ping", func(t *testing.T) {
		router := restServer(context.Background(), addr)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/ping", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "pong", w.Body.String())
	})

	t.Run("grpc", func(t *testing.T) {
		name := "테스트네임"
		router := restServer(context.Background(), addr)

		b, er := json.Marshal(map[string]interface{}{"name": name})
		assert.NoError(t, er)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/v1/post", bytes.NewReader(b))
		router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)

		d := map[string]string{}
		er = json.Unmarshal(w.Body.Bytes(), &d)
		assert.NoError(t, er)
		assert.Equal(t, name, d["message"])
	})

}
