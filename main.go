package main

import (
	"fmt"
	"net/http"
	"unit-converter/server"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

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

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
