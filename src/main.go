package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
	"google.golang.org/api/sheets/v4"
)

var service *sheets.Service
var storedSessions []*SessionInfo

func UpdateStoredSessions() {
	ss, err := ReadSheets(ReadOptions{
		SheetId: os.Getenv("SHEET_ID"),
		Service: service,
		Range:   "A2:G",
	})
	if err == nil {
		storedSessions = ss
	}
}

func main() {
	godotenv.Load()
	service, _ = InitSheets(os.Getenv("SHEETS_SERVICE_ACCOUNT_KEY"), context.Background())

	UpdateStoredSessions()

	// Bot initialization code adapted from
	// https://github.com/slack-go/slack/blob/master/examples/socketmode/socketmode.go
	appToken := os.Getenv("SLACK_APP_TOKEN")

	if appToken == "" {
		fmt.Fprintf(os.Stderr, "SLACK_APP_TOKEN must be set.")
		os.Exit(1)
	}

	if !strings.HasPrefix(appToken, "xapp-") {
		fmt.Fprintf(os.Stderr, "SLACK_APP_TOKEN is not a valid token")
		os.Exit(1)
	}

	botToken := os.Getenv("SLACK_BOT_TOKEN")
	if botToken == "" {
		fmt.Fprintf(os.Stderr, "SLACK_BOT_TOKEN must be set")
		os.Exit(1)
	}

	if !strings.HasPrefix(botToken, "xoxb-") {
		fmt.Fprintf(os.Stderr, "SLACK_BOT_TOKEN is not a valid token")
		os.Exit(1)
	}

	api := slack.New(
		botToken,
		slack.OptionDebug(true),
		slack.OptionLog(log.New(os.Stdout, "api: ", log.Lshortfile|log.LstdFlags)),
		slack.OptionAppLevelToken(appToken),
	)

	client := socketmode.New(
		api,
		socketmode.OptionDebug(true),
		socketmode.OptionLog(log.New(os.Stdout, "socket: ", log.Lshortfile|log.LstdFlags)),
	)

	go func() {
		sessions := make(map[string]*SessionInfo)

		for evt := range client.Events {
			switch evt.Type {
			case socketmode.EventTypeConnecting:
				fmt.Println("Connecting to Slack in Socket Mode")
			case socketmode.EventTypeConnectionError:
				fmt.Println("Failed to connect to Slack in Socket Mode")
			case socketmode.EventTypeConnected:
				fmt.Println("Connected to Slack in Socket Mode")
			case socketmode.EventTypeEventsAPI:
				handleEventsAPIEvent(client, evt, sessions)
			case socketmode.EventTypeInteractive:
				handleInteractiveEvent(client, evt, sessions)
			case socketmode.EventTypeSlashCommand:
				handleSlashCommandEvent(client, evt, sessions)
			default:
				fmt.Fprintf(os.Stderr, "Unexpected event type recieved: %s\n", evt.Type)
			}
		}
	}()

	client.Run()
}

func handleEventsAPIEvent(client *socketmode.Client, event socketmode.Event, sessions map[string]*SessionInfo) error {
	client.Ack(*event.Request)
	return errors.New("not implemented")
}

func handleInteractiveEvent(client *socketmode.Client, event socketmode.Event, sessions map[string]*SessionInfo) error {
	callback, ok := event.Data.(slack.InteractionCallback)
	if !ok {
		return errors.New("type assertion failed")
	}

	userId := callback.User.ID

	// Delete the ongoing session if the user closed the modal
	if callback.Type == slack.InteractionTypeViewClosed {
		delete(sessions, userId)
		client.Ack(*event.Request)
		return nil
	}

	if callback.View.CallbackID == "team-information-callback" {
		// Get the submitted values and put them in the session object for this user
		selectedMachine := callback.View.State.Values["machine_input_block"]["machine_select"].SelectedOption.Value
		selectedUsers := callback.View.State.Values["users_input_block"]["users_select"].SelectedUsers
		session := sessions[userId]
		session.GroupIds = selectedUsers
		session.Machine = selectedMachine
		sessions[userId] = session

		updateView(client, event, MachineInformation(session))
	} else if callback.View.CallbackID == "machine-information-callback" {
		machineReason := callback.View.State.Values["machine_purpose_input_block"]["machine_purpose"].Value
		machineDuration := callback.View.State.Values["duration_input_block"]["duration_select"].SelectedOption.Value
		session := sessions[userId]
		session.Reason = machineReason
		session.Duration = DurationToMins(machineDuration)

		sessions[userId] = session

		updateView(client, event, TimeInformation(session, storedSessions))
	} else if callback.View.CallbackID == "time-information-callback" {
		timeSlot := callback.View.State.Values["time_input_block"]["time_select"].SelectedOption.Value
		session := sessions[userId]
		session.Time, _ = time.Parse(time.RFC3339, timeSlot)
		sessions[userId] = session

		client.Ack(*event.Request)

		WriteSheets(WriteOptions{
			SheetId: os.Getenv("SHEET_ID"),
			Service: service,
			Range:   "A2:G",
			Session: session,
		})

	} else {
		// Required Ack so slack knows the bot didn't die
		client.Ack(*event.Request)
	}

	go UpdateStoredSessions()

	// No error, return nil.
	return nil
}

func updateView(client *socketmode.Client, event socketmode.Event, view slack.ModalViewRequest) {
	client.Ack(*event.Request, struct {
		ResponseAction string      `json:"response_action"`
		View           interface{} `json:"view"`
	}{
		ResponseAction: "update",
		View:           view,
	})
}

func handleSlashCommandEvent(client *socketmode.Client, event socketmode.Event, sessions map[string]*SessionInfo) error {
	command, ok := event.Data.(slack.SlashCommand)
	if !ok {
		return errors.New("type assertion failed")
	}

	if command.Command == "/test" {

		// Slash command will always open the first modal.
		view := TeamInformation()
		resp, err := client.OpenView(command.TriggerID, view)
		if err != nil {
			return err
		}

		// Store a session keyed on the userId.
		userId := command.UserID
		userName := command.UserName
		viewId := resp.View.ID
		session := &SessionInfo{}
		session.UserId = userId
		session.UserName = userName
		session.ViewId = viewId
		sessions[userId] = session

		client.Ack(*event.Request)
	} else if command.Command == "/refresh" {
		UpdateStoredSessions()
		client.Ack(*event.Request)
	}
	return nil
}
