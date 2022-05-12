package server

import (
	"net/http"

	"github.com/missingsemi/capstone/frontend"
)

func RegisterRoutes() {
	http.HandleFunc("/api/schedule", CorsMiddleware(HandleSchedule))
	http.HandleFunc("/api/machines", CorsMiddleware(HandleMachines))
	http.Handle("/", http.FileServer(http.FS(frontend.Frontend)))
}
