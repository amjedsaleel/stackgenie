package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	contentPath := flag.String("path", "./static", "Static content absolute path,")
	listenServer := flag.String("listen", "localhost:8080", "Server listening to,")

	flag.Parse()
	fmt.Println("\n  Simple HTTP server running on http://"+*listenServer, "\n  Static content directory", *contentPath)

	http.Handle("/", http.FileServer(http.Dir(*contentPath)))
	log.Fatal(http.ListenAndServe(*listenServer, nil))

}
