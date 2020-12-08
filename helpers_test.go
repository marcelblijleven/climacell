package climacell

import (
	"net/http"
	"net/http/httptest"
)

func setupTestServer(handlerFunc func(w http.ResponseWriter, r *http.Request)) (*httptest.Server, func()) {
	srv := httptest.NewServer(http.HandlerFunc(handlerFunc))
	backupBaseURL := BaseURL
	BaseURL = srv.URL

	closeFunc := func() {
		srv.Close()
		BaseURL = backupBaseURL
	}

	return srv, closeFunc
}
