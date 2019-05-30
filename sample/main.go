package main

import (
	"fmt"
	"net/http"

	"github.com/yasukotelin/go-openurl"
)

const (
	indexPath = ""
	port      = ":8080"
)

func main() {
	routing()
	listen()
}

func getURL() string {
	return fmt.Sprintf("http://localhost%s", port)
}

func routing() {
	http.Handle("/", http.FileServer(http.Dir("./")))
}

func listen() {
	fmt.Println("server started.")
	fmt.Printf("listen on the localhost%s\n", port)
	fmt.Println()
	fmt.Println("stop server pressed Ctrl+C.")
	if err := openurl.OpenWithBrowser(getURL()); err != nil {
		fmt.Println(err)
	}
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println(err)
		return
	}
}
