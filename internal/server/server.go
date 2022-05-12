package server

import (
	"fmt"
	"log"
	"mime"
	"net/http"
	"strings"
	"sync"

	"github.com/missingsemi/capstone/internal/slackutil"
)

func Start(port int, appToken string, botToken string, wg *sync.WaitGroup) {
	defer wg.Done()

	if appToken == "" {
		panic("SLACK_APP_TOKEN must be set.")
	}

	if !strings.HasPrefix(appToken, "xapp-") {
		panic("SLACK_APP_TOKEN is not a valid token")
	}

	slackutil.Client(appToken, botToken)

	RegisterRoutes()

	mime.AddExtensionType(".js", "text/javascript")
	mime.AddExtensionType(".html", "text/html")
	mime.AddExtensionType(".htm", "text/html")

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
