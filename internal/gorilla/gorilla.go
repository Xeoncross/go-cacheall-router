package gorilla

import (
	"net/http"
	"net/url"
)

// Handler returns gorilla/mux handler
func Handler(dir http.FileSystem, api http.HandlerFunc) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api", http.HandlerFunc(api))
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(dir)))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Overwrite URL???
		// r.URL = &url.URL{Path: "index.html"}

		u, _ := url.Parse("index.html")
		r.URL = u

		http.FileServer(dir).ServeHTTP(w, r)

		// Pull from memory object as io.Reader??
		// w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// http.ServeContent(w, r, "index.html", time.Time{}, gocacheallrouter.MemoryHTTPFilesystem.Open('..))
	})
	return mux
}
