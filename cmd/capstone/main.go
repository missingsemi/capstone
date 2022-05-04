package main

import (
	"flag"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/missingsemi/capstone/pkg/bot"
	"github.com/missingsemi/capstone/pkg/database"
	"github.com/missingsemi/capstone/pkg/server"
)

func main() {
	godotenv.Load()
	startBot := flag.Bool("bot", true, "If false, doesn't launch the slack bot.")
	hostServer := flag.Bool("host", true, "If false, doesn't launch the web interface.")
	port := flag.Int("port", 8080, "Sets the port the web interface is hosted on.")

	flag.Parse()

	database.DbInit()
	defer database.DbDeinit()

	wg := sync.WaitGroup{}

	if *startBot {
		appToken := os.Getenv("SLACK_APP_TOKEN")
		botToken := os.Getenv("SLACK_BOT_TOKEN")
		wg.Add(1)
		go bot.Start(appToken, botToken, &wg)
	}

	if *hostServer {
		wg.Add(1)
		go server.Start(*port, &wg)
	}

	wg.Wait()
}
