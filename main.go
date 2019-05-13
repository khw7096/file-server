package main

import (
	"flag"
	"fmt"
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
	ip, err := externalIP()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("service start : http://%s%s\n", ip, addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
