package singlepageapp

import (
	"net/http"
	"os"
)

// FileSystem that routes all requests for invalid paths to index.html
func FileSystem(fs http.FileSystem) http.FileSystem {
	return &spaFileSystem{fs}
}

type spaFileSystem struct {
	root http.FileSystem
}

// Open file or default to index.html
func (fs *spaFileSystem) Open(name string) (http.File, error) {
	f, err := fs.root.Open(name)
	if os.IsNotExist(err) {
		return fs.root.Open("index.html")
	}
	return f, err
}
