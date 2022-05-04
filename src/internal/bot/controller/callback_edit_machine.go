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

func CallbackEditMachine(client *socketmode.Client, event socketmode.Event) error {
	callback, ok := event.Data.(slack.InteractionCallback)
	if !ok {
		return errors.New("type assertion failed")
	}

	idInput := callback.View.State.Values["id_input_block"]["id_input"].Value
	nameInput := callback.View.State.Values["name_input_block"]["name_input"].Value
	titleNameInput := callback.View.State.Values["titlename_input_block"]["titlename_input"].Value
	countInput := callback.View.State.Values["count_input_block"]["count_input"].Value

	machineId := callback.View.PrivateMetadata
	machine, _ := database.GetMachineById(machineId)

	if idInput == "" {
		idInput = machine.Id
	}

	if nameInput == "" {
		nameInput = machine.Name
	}

	if titleNameInput == "" {
		titleNameInput = machine.TitleName
	}

	var parsedCount int64
	if countInput == "" {
		parsedCount = int64(machine.Count)
	} else {
		var err error
		parsedCount, err = strconv.ParseInt(callback.View.State.Values["count_input_block"]["count_input"].Value, 10, 32)
		if err != nil {

			util.ErrorView(
				client,
				event,
				map[string]string{
					"count_input_block": "Count must be an integer.",
				},
			)

			return err
		}
	}

	database.ModifyMachine(machine.Id, model.Machine{
		Id:        idInput,
		Name:      nameInput,
		TitleName: titleNameInput,
		Count:     int(parsedCount),
	})

	client.Ack(*event.Request)

	machines, _ := database.GetMachines()
	util.UpdateView2(client, callback.View.PreviousViewID, view.AdminMachinesAvailableMachines(machines))

	return nil
}
