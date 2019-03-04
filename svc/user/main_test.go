package main

/*
func TestRoute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	post := postmock.NewMockPostServiceClient(ctrl)
	cc := &connectors{post}

	initTest(t)
	gs := startMockGrpcServer(cc)
	defer gs.Stop()

	t.Run("ping", func(t *testing.T) {
		router := mockRouter()

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/ping", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "pong", w.Body.String())
	})

	t.Run("grpc", func(t *testing.T) {
		content := "테스트user"
		router := mockRouter()

		s := fmt.Sprintf(`{"content": "%s"}`, content)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/tokens", strings.NewReader(s))
		router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)

		d := struct {
			Token struct {
				AccessToken string
				ExpiresIn   string
			}
		}{}
		er := json.Unmarshal(w.Body.Bytes(), &d)
		assert.NoError(t, er)

		assert.NotEmpty(t, d.Token.AccessToken)
		assert.NotEmpty(t, d.Token.ExpiresIn)
	})

}

*/
