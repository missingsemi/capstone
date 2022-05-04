package server

import (
	"fmt"
	"log"
	"mime"
	"net/http"
)

func Start(port int) {
	RegisterRoutes()

	mime.AddExtensionType(".js", "text/javascript")
	mime.AddExtensionType(".html", "text/html")
	mime.AddExtensionType(".htm", "text/html")

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
