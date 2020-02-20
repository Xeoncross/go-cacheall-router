package gorilla

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"time"
)

// Handler returns gorilla/mux handler
func Handler(dir http.FileSystem, api http.HandlerFunc) http.Handler {

	f, err := dir.Open("index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b := &bytes.Buffer{}
	io.Copy(b, f)
	err = f.Close()

	b2 := bytes.NewReader(b.Bytes())

	mux := http.NewServeMux()
	mux.HandleFunc("/api", http.HandlerFunc(api))
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(dir)))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Overwrite URL???
		// r.URL = &url.URL{Path: "index.html"}
		// u, _ := url.Parse("/index.html")
		// r.URL = u
		// http.FileServer(dir).ServeHTTP(w, r)

		// Pull from memory object as io.Reader??
		// w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// http.ServeContent(w, r, "index.html", time.Time{}, gocacheallrouter.MemoryHTTPFilesystem.Open('..))

		// Just buffer the contents once and send every time?
		http.ServeContent(w, r, "index.html", time.Time{}, b2)

	})
	return mux
}
