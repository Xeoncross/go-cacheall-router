package gorilla

import (
	"net/http"
)

// Handler returns gorilla/mux handler
// func Handler(dir http.FileSystem, api http.HandlerFunc) http.Handler {

// 	f, err := dir.Open("index.html")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()

// 	b, err := ioutil.ReadAll(f)

// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/api", http.HandlerFunc(api))
// 	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(dir)))
// 	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

// 		// Overwrite URL???
// 		// r.URL = &url.URL{Path: "index.html"}
// 		// u, _ := url.Parse("/index.html")
// 		// r.URL = u
// 		// http.FileServer(dir).ServeHTTP(w, r)

// 		// Pull from memory object as io.Reader??
// 		// w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 		// http.ServeContent(w, r, "index.html", time.Time{}, gocacheallrouter.MemoryHTTPFilesystem.Open('..))

// 		// One reader per client because it tracks user state
// 		readseeker := bytes.NewReader(b)

// 		// Just buffer the contents once and send every time?
// 		http.ServeContent(w, r, "index.html", time.Time{}, readseeker)

// 	})
// 	return mux
// }

// Handler returns gorilla/mux handler
func Handler(dir http.FileSystem, api http.HandlerFunc) http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/api", http.HandlerFunc(api))
	mux.Handle("/", http.FileServer(dir))
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(dir)))
	return mux
}
