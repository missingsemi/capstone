package controller

import (
	"errors"
	"strconv"

	"github.com/missingsemi/capstone/internal/bot/util"
	"github.com/missingsemi/capstone/internal/bot/view"
	"github.com/missingsemi/capstone/pkg/database"
	"github.com/missingsemi/capstone/pkg/model"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func CallbackCreateMachine(client *socketmode.Client, event socketmode.Event) error {
	callback, ok := event.Data.(slack.InteractionCallback)
	if !ok {
		return errors.New("type assertion failed")
	}

	idInput := callback.View.State.Values["id_input_block"]["id_input"].Value
	nameInput := callback.View.State.Values["name_input_block"]["name_input"].Value
	titleNameInput := callback.View.State.Values["titlename_input_block"]["titlename_input"].Value
	countInput, err := strconv.ParseInt(callback.View.State.Values["count_input_block"]["count_input"].Value, 10, 32)

	if err != nil {
		errors := make(map[string]string, 1)
		errors["count_input_block"] = "Count must be an integer"

		util.ErrorView(client, event, errors)

		return err
	}

	database.CreateMachine(model.Machine{
		Id:        idInput,
		Name:      nameInput,
		TitleName: titleNameInput,
		Count:     int(countInput),
	})

	client.Ack(*event.Request)

	machines, _ := database.GetMachines()
	util.UpdateView2(client, callback.View.PreviousViewID, view.AdminMachinesAvailableMachines(machines))

	return nil
}
