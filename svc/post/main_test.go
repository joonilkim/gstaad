package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	assert "github.com/stretchr/testify/require"
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
		content := "테스트"
		router := restServer(context.Background(), addr)

		s := fmt.Sprintf(`{"content": "%s"}`, content)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/posts", strings.NewReader(s))
		router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)

		d := map[string]interface{}{}
		er := json.Unmarshal(w.Body.Bytes(), &d)
		assert.NoError(t, er)
		assert.Equal(t, true, d["result"])
	})

}
