package gocacheallrouter

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRawHandlers(t *testing.T) {

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
			handler:  http.FileServer(MemoryHTTPFilesystem),
			response: IndexResponse,
		},
		{
			desc:     "missing",
			path:     "/foo/bar",
			handler:  http.FileServer(MemoryHTTPFilesystem),
			response: IndexResponse,
		},
		{
			desc:     "app.js",
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

// func TestHandlers(t *testing.T) {

// 	testCases := []struct {
// 		desc     string
// 		path     string
// 		handler  http.Handler
// 		response string
// 	}{
// 		{
// 			desc:     "api",
// 			path:     "/api",
// 			handler:  http.HandlerFunc(APIHandler),
// 			response: APIResponse,
// 		},
// 		{
// 			desc:     "index",
// 			path:     "/",
// 			handler:  http.FileServer(MemoryHTTPFilesystem),
// 			response: IndexResponse,
// 		},
// 		{
// 			desc:     "index",
// 			path:     "/foo/bar",
// 			handler:  http.FileServer(MemoryHTTPFilesystem),
// 			response: IndexResponse,
// 		},
// 		{
// 			desc:     "api",
// 			path:     "/js/app.js",
// 			handler:  http.FileServer(MemoryHTTPFilesystem),
// 			response: JavascriptResponse,
// 		},
// 	}

// 	tests := []struct {
// 		desc    string
// 		handler http.Handler
// 	}{
// 		{
// 			desc:    "raw",
// 			handler: nil,
// 		},
// 		{
// 			desc:    "gorilla",
// 			handler: gorilla.Handler(),
// 		},
// 	}

// 	for _, test := range tests {
// 		for _, tc := range testCases {
// 			t.Run(tc.desc, func(t *testing.T) {

// 				req, err := http.NewRequest("GET", tc.path, nil)
// 				if err != nil {
// 					t.Fatal(err)
// 				}

// 				rr := httptest.NewRecorder()

// 				test.handler.ServeHTTP(rr, req)

// 				if rr.Body.String() != tc.response {
// 					t.Errorf("got: %q\nwant: %q\n", rr.Body.String(), tc.response)
// 				}
// 			})
// 		}
// 	}
// }
