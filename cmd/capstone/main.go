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
	// startBot := flag.Bool("bot", true, "Whether or not to launch the slack bot.")
	// hostServer := flag.Bool("website", true, "Whether or not to launch the web interface.")
	// port := flag.Int("port", 8080, "The port to host the web interface on.")
	// dbfile := flag.String("dbfile", "schedule.db", "The path to the sqlite file to store session data in.")

	startBot := flag.Bool("bot", true, "launch the slack bot")
	hostServer := flag.Bool("site", true, "host the website and api")
	port := flag.Int("port", 8080, "port to host the website and api on")
	dbfile := flag.String("dbfile", "schedule.db", "path to the sqlite database file")
	init := flag.Bool("init", false, "create and configure a new database file")
	appToken := flag.String("apptoken", "", "slack app token. flag takes precedence over SLACK_APP_TOKEN environment variable")
	botToken := flag.String("bottoken", "", "slack bot token. flag takes precedence over SLACK_BOT_TOKEN environment variable")
	admin := flag.String("admin", "", "add a new admin")

	flag.Parse()

	database.DbInit(*dbfile, *init)
	defer database.DbDeinit()

	if *admin != "" {
		database.CreateAdmin(*admin)
	}

	wg := sync.WaitGroup{}

	if *appToken == "" {
		*appToken = os.Getenv("SLACK_APP_TOKEN")
	}

	if *botToken == "" {
		*botToken = os.Getenv("SLACK_BOT_TOKEN")
	}

	if *startBot {
		wg.Add(1)
		go bot.Start(*appToken, *botToken, &wg)
	}

	if *hostServer {
		wg.Add(1)
		go server.Start(*port, *appToken, *botToken, &wg)
	}

	wg.Wait()
}
