package bot

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/missingsemi/capstone/internal/bot/controller"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func Start(appToken string, botToken string, wg *sync.WaitGroup) {
	defer wg.Done()

	controller.RegisterCommandHandler("/schedule", controller.CommandSchedule)
	controller.RegisterCallbackHandler("user_schedule-created_sessions-callback", controller.CallbackCreatedSessions)
	controller.RegisterCallbackHandler("user_schedule-create_session_1-callback", controller.CallbackCreateSession1)
	controller.RegisterCallbackHandler("user_schedule-create_session_2-callback", controller.CallbackCreateSession2)

	controller.RegisterCommandHandler("/machines", controller.CommandMachines)
	controller.RegisterCallbackHandler("admin_machines-available_machines-callback", controller.CallbackAvailableMachines)
	controller.RegisterCallbackHandler("admin_machines-edit_machine-callback", controller.CallbackEditMachine)
	controller.RegisterCallbackHandler("admin_machines-create_machine-callback", controller.CallbackCreateMachine)

	// Bot initialization code adapted from
	// https://github.com/slack-go/slack/blob/master/examples/socketmode/socketmode.go

	if appToken == "" {
		panic("SLACK_APP_TOKEN must be set.")
	}

	if !strings.HasPrefix(appToken, "xapp-") {
		panic("SLACK_APP_TOKEN is not a valid token")
	}

	if botToken == "" {
		panic("SLACK_BOT_TOKEN must be set")
	}

	if !strings.HasPrefix(botToken, "xoxb-") {
		panic("SLACK_BOT_TOKEN is not a valid token")
	}

	api := slack.New(
		botToken,
		slack.OptionDebug(true),
		//slack.OptionLog(log.New(os.Stdout, "api: ", log.Lshortfile|log.LstdFlags)),
		slack.OptionAppLevelToken(appToken),
	)

	client := socketmode.New(
		api,
		socketmode.OptionDebug(true),
		//socketmode.OptionLog(log.New(os.Stdout, "socket: ", log.Lshortfile|log.LstdFlags)),
	)

	go controller.Notify(api)

	go func() {
		for evt := range client.Events {
			switch evt.Type {
			case socketmode.EventTypeConnecting:
				fmt.Println("Connecting to Slack in Socket Mode")
			case socketmode.EventTypeConnectionError:
				fmt.Println("Failed to connect to Slack in Socket Mode")
			case socketmode.EventTypeConnected:
				fmt.Println("Connected to Slack in Socket Mode")
			case socketmode.EventTypeEventsAPI:
				client.Ack(*evt.Request)
			case socketmode.EventTypeInteractive:
				callback, ok := evt.Data.(slack.InteractionCallback)
				if !ok {
					client.Ack(*evt.Request)
				}
				controller.CallCallbackHandler(callback.View.CallbackID, client, evt)
			case socketmode.EventTypeSlashCommand:
				cmd, ok := evt.Data.(slack.SlashCommand)
				if !ok {
					client.Ack(*evt.Request)
				}
				controller.CallCommandHandler(cmd.Command, client, evt)
			default:
				fmt.Fprintf(os.Stderr, "Unexpected event type recieved: %s\n", evt.Type)
			}
		}
	}()

	client.Run()
}
