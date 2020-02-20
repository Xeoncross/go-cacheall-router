package gorilla

import (
	"net/http"
	"net/url"

	gocacheallrouter "github.com/xeoncross/go-cacheall-router"
)

// Handler returns gorilla/mux handler
func Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api", http.HandlerFunc(gocacheallrouter.APIHandler))
	mux.Handle("/build/", http.StripPrefix("/build/", http.FileServer(gocacheallrouter.MemoryHTTPFilesystem)))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Overwrite URL???
		r.URL = &url.URL{Path: "index.html"}
		http.FileServer(gocacheallrouter.MemoryHTTPFilesystem).ServeHTTP(w, r)

		// Pull from memory object as io.Reader??
		// w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// http.ServeContent(w, r, "index.html", time.Time{}, gocacheallrouter.MemoryHTTPFilesystem.Open('..))
	})
	return mux
}
