package gocacheallrouter

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers(t *testing.T) {

	testCases := []struct {
		desc     string
		path     string
		handler  http.Handler
		response string
	}{
		{
			desc:     "api",
			path:     "/",
			handler:  http.HandlerFunc(APIHandler),
			response: APIResponse,
		},
		{
			desc:     "index",
			path:     "/",
			handler:  http.FileServer(MemoryHTTPFilesystem),
			response: IndexResponse,
		},
		{
			desc:     "api",
			path:     "/js/app.js",
			handler:  http.FileServer(MemoryHTTPFilesystem),
			response: JavascriptResponse,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {

			req, err := http.NewRequest("GET", tc.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()

			tc.handler.ServeHTTP(rr, req)

			if rr.Body.String() != tc.response {
				t.Errorf("got: %q\nwant: %q\n", rr.Body.String(), tc.response)
			}
		})
	}

}
