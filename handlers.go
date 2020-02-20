package gocacheallrouter

import (
	"log"
	"net/http"

	"github.com/spf13/afero"
)

// IndexResponse body
const IndexResponse = "<script src='/js/app.js'></script>"

// JavascriptResponse body
const JavascriptResponse = "document.write('JS loaded')"

// APIResponse body
const APIResponse = "api"

// MemoryHTTPFilesystem to pretend we are reading from a SPA directory
var MemoryHTTPFilesystem *afero.HttpFs

// IndexHandler would serve the index.html page built by the SPA
// var IndexHandler = func(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte(IndexResponse))
// }

// // JavascriptHandler would serve the index.html page built by the SPA
// var JavascriptHandler = func(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte(IndexResponse))
// }

// APIHandler would represent a response from the Go API
var APIHandler = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(APIResponse))
}

func init() {
	var fs = afero.NewMemMapFs()

	file, err := fs.Create("/index.html")
	if err != nil {
		log.Fatal(err)
	}
	file.Write([]byte(IndexResponse))
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	file, err = fs.Create("/js/app.js")
	if err != nil {
		log.Fatal(err)
	}
	file.Write([]byte(JavascriptResponse))
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	MemoryHTTPFilesystem = afero.NewHttpFs(fs)
	// fileserver := http.FileServer(httpFs.Dir("templates"))
	// http.Handle("/", fileserver)
}
