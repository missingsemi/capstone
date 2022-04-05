package db

import (
	"context"

	"github.com/missingsemi/capstone/model"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type WriteOptions struct {
	SheetId string
	Range   string
	Session *model.ScheduleAddSession
	Service *sheets.Service
}

type ReadOptions struct {
	SheetId string
	Range   string
	Limit   int
	Service *sheets.Service
}

func InitSheets(apiKey string, ctx context.Context) (*sheets.Service, error) {
	return sheets.NewService(
		ctx,
		option.WithScopes(sheets.SpreadsheetsScope),
		//option.WithAPIKey(apiKey),
		option.WithCredentialsFile(apiKey),
	)
}

func WriteSheets(options WriteOptions) error {
	values := sheets.ValueRange{
		Values: [][]interface{}{SessionToRow(options.Session)},
	}

	_, err := options.Service.Spreadsheets.Values.Append(
		options.SheetId,
		options.Range,
		&values,
	).ValueInputOption("USER_ENTERED").Do()

	if err != nil {
		return err
	}

	return nil
}

func ReadSheets(options ReadOptions) ([]*model.ScheduleAddSession, error) {
	resp, err := options.Service.Spreadsheets.Values.Get(
		options.SheetId,
		options.Range,
	).Do()
	if err != nil {
		return nil, err
	}

	result := make([]*model.ScheduleAddSession, 0)
	for _, row := range resp.Values {
		session := SessionFromRow(row)
		if session != nil {
			result = append(result, session)
		}
	}

	return result, nil
}

type FilterOptions struct {
	UserId  string
	Machine string
}

func FilterSessions(sessions []*model.ScheduleAddSession, options FilterOptions) []*model.ScheduleAddSession {
	result := make([]*model.ScheduleAddSession, 0)
	for _, session := range sessions {
		if options.UserId != "" && options.UserId != session.UserId {
			continue
		}
		if options.Machine != "" && options.Machine != session.Machine {
			continue
		}
		result = append(result, session)
	}

	return result
}
