package gocacheallrouter

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/xeoncross/go-cacheall-router/internal/gorilla"
	"github.com/xeoncross/go-cacheall-router/internal/singlepageapp"
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
			path:     "/api",
			handler:  http.HandlerFunc(APIHandler),
			response: APIResponse,
		},
		{
			desc:     "index",
			path:     "/",
			handler:  http.FileServer(singlepageapp.FileSystem(MemoryHTTPFilesystem)),
			response: IndexResponse,
		},
		{
			desc:     "javascript",
			path:     "/js/app.js",
			handler:  http.FileServer(singlepageapp.FileSystem(MemoryHTTPFilesystem)),
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

	//
	// Now test libraries
	//

	handlers := []struct {
		desc    string
		handler http.Handler
	}{
		{
			desc:    "gorilla/mux/",
			handler: gorilla.Handler(singlepageapp.FileSystem(MemoryHTTPFilesystem), APIHandler),
		},
	}

	for _, tc := range testCases {
		for _, handler := range handlers {

			t.Run(handler.desc+tc.desc, func(t *testing.T) {
				req, err := http.NewRequest("GET", tc.path, nil)
				if err != nil {
					t.Fatal(err)
				}

				rr := httptest.NewRecorder()

				handler.handler.ServeHTTP(rr, req)

				if rr.Body.String() != tc.response {
					t.Errorf("got: %q\nwant: %q\n", rr.Body.String(), tc.response)
				}
			})
		}
	}

}
