package server

import (
	"sync"

	"github.com/missingsemi/capstone/internal/server"
)

func Start(port int, wg *sync.WaitGroup) {
	server.Start(port)
	defer wg.Done()
}
