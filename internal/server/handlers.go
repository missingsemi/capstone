package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/missingsemi/capstone/pkg/database"
)

func HandleSchedule(w http.ResponseWriter, req *http.Request) {
	// parse date from query params
	// if absent, default to time.Now()
	// otherwise, 400 if theres a parse error
	queryDate := req.URL.Query().Get("date")
	var date time.Time
	if queryDate == "" {
		date = time.Now()
	} else {
		var err error
		date, err = time.Parse(time.RFC3339, queryDate)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	// Search from start of day to end of day
	year, month, day := date.Date()
	date = time.Date(year, month, day, 0, 0, 0, 0, date.Location())
	endDate := time.Date(year, month, day+1, 0, 0, 0, 0, date.Location())

	// fetch from db and 500 on fail
	sessions, err := database.GetSessionsBetweenTimes(date, endDate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// generate body and 500 on fail
	j, err := json.Marshal(sessions)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// write body + mime and default 200
	w.Header().Add("Content-Type", "application/json")
	w.Write(j)
	return
}

func HandleMachines(w http.ResponseWriter, req *http.Request) {
	// fetch from db and 500 on fail
	machines, err := database.GetMachines()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// generate body and 500 on fail
	j, err := json.Marshal(machines)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// write body + mime and default 200
	w.Header().Add("Content-Type", "application/json")
	w.Write(j)
	return
}
