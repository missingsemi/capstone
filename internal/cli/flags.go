package cli

import "flag"

func ParseFlags(config *Config) {
	// godotenv.Load()
	// startBot := flag.Bool("bot", true, "Whether or not to launch the slack bot.")
	// hostServer := flag.Bool("website", true, "Whether or not to launch the web interface.")
	// port := flag.Int("port", 8080, "The port to host the web interface on.")

	// flag.Parse()

	host := flag.Bool("host", true, "Sets if the web server is launched.")
	port := flag.Int("port", 8080, "Sets the port the web server is launched on.")
	frontend := flag.Bool("frontend", false, "Controls if frontend routes are registered.")
	bot := flag.Bool("bot", true, "Sets if the slack bot is launched.")
	appToken := flag.String("app-token", "", "Slack app token. If not set, the SLACK_APP_TOKEN environment variable is checked.")
	botToken := flag.String("bot-token", "", "Slack bot token. If not set, the SLACK_BOT_TOKEN environment variable is checked.")
	databaseLocation := flag.String("database", "./schedule.db", "Database file to connect to.")

	flag.Parse()

	config.Host = *host
	config.Port = *port
	config.Frontend = *frontend
	config.Bot = *bot
	config.AppToken = *appToken
	config.BotToken = *botToken
	config.DatabaseLocation = *databaseLocation
}
