package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var (
		addr string
		dir  string
	)
	flag.StringVar(&addr, "addr", ":80", "address to bind")
	flag.StringVar(&dir, "dir", ".", "directory to serve")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(dir)))

	err := http.ListenAndServe(addr, nil)
	log.Fatal(err)
}
