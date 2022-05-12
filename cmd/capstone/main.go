package main

import (
	"flag"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/missingsemi/capstone/internal/bot"
	"github.com/missingsemi/capstone/internal/database"
	"github.com/missingsemi/capstone/internal/server"
)

func main() {
	godotenv.Load()
	startBot := flag.Bool("bot", true, "Whether or not to launch the slack bot.")
	hostServer := flag.Bool("website", true, "Whether or not to launch the web interface.")
	port := flag.Int("port", 8080, "The port to host the web interface on.")
	dbfile := flag.String("dbfile", "schedule.db", "The path to the sqlite file to store session data in.")

	flag.Parse()

	database.DbInit(*dbfile)
	defer database.DbDeinit()

	wg := sync.WaitGroup{}

	appToken := os.Getenv("SLACK_APP_TOKEN")
	botToken := os.Getenv("SLACK_BOT_TOKEN")

	if *startBot {
		wg.Add(1)
		go bot.Start(appToken, botToken, &wg)
	}

	if *hostServer {
		wg.Add(1)
		go server.Start(*port, appToken, botToken, &wg)
	}

	wg.Wait()
}
