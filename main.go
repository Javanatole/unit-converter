package main

import (
	"fmt"
	"net/http"
	"unit-converter/server"
)

func main() {
	// provide all static html files
	http.Handle("/", http.FileServer(http.Dir("static")))

	// init all http routes
	server.InitHttpRoutes()

	// show ready to server and start server
	fmt.Println("Ready to serve unit converter on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Error while serving unit converter")
	}
}
