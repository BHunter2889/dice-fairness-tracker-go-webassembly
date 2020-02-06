// +build !wasm

package main

//go:generate vugugen .

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/vugu/vugu/simplehttp"
)

func main() {
	port := os.Getenv("PORT")
	log.Printf("Provided Port: %q", port)
	if port == "" {
		port = "127.0.0.1:8877"
	}

	dev := flag.Bool("dev", false, "Enable development features")
	dir := flag.String("dir", ".", "Project directory")
	httpl := flag.String("http", port, "Listen for HTTP on this host:port")
	flag.Parse()
	wd, _ := filepath.Abs(*dir)
	os.Chdir(wd)
	log.Printf("Starting HTTP Server at %q", *httpl)
	h := simplehttp.New(wd, *dev)
	// import CSS / fonts
	simplehttp.DefaultStaticData["CSSFiles"] = []string{ "static/css/fonts.css", "static/css/index.css" }
	log.Fatal(http.ListenAndServe(*httpl, h))
}