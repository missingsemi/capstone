package cli

type Config struct {
	// Controls if the web server is launched
	Host bool
	// Port to host the web server on
	Port int
	// Controls if frontend routes are registered or not
	Frontend bool
	// Controls if the slack bot is launched
	Bot bool
	// Slack app token
	AppToken string
	// Slack bot token
	BotToken string
	// Location of the database file
	DatabaseLocation string
	// Explicit flag to call Setup
	Init bool
}
