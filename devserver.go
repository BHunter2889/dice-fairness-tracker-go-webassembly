// +build ignore

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/vugu/vugu/simplehttp"
)

func main() {
	wd, _ := os.Getwd()
	l := "127.0.0.1:8844"
	log.Printf("Starting HTTP Server at %q", l)
	h := simplehttp.New(wd, true)
	// import CSS / fonts
	simplehttp.DefaultStaticData["CSSFiles"] = []string{ "static/css/fonts.css", "static/css/index.css" }
	log.Fatal(http.ListenAndServe(l, h))
}