package server

import "net/http"

func RegisterRoutes() {
	http.HandleFunc("/api/schedule", CorsMiddleware(HandleSchedule))
	http.HandleFunc("/api/machines", CorsMiddleware(HandleMachines))
	http.Handle("/", http.FileServer(http.Dir("./frontend/")))
}
