package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/missingsemi/capstone/internal/database"
	"github.com/missingsemi/capstone/internal/slackutil"
)

type scheduleBody struct {
	Id         int      `json:"id"`
	Username   string   `json:"username"`
	GroupNames []string `json:"groupNames"`
	Machine    string   `json:"machine"`
	Reason     string   `json:"reason"`
	Duration   int      `json:"duration"`
	Time       string   `json:"time"`
}

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

	cachedUsernames := make(map[string]string, 0)

	body := make([]scheduleBody, 0, len(sessions))
	_, apiClient := slackutil.Client("", "")
	for _, session := range sessions {
		if _, ok := cachedUsernames[session.UserId]; !ok {
			username, _ := slackutil.GetUsername(apiClient, session.UserId)
			cachedUsernames[session.UserId] = username
		}

		groupNames := make([]string, 0, len(session.GroupIds))
		for _, groupId := range session.GroupIds {
			if username, ok := cachedUsernames[groupId]; !ok {
				username, _ = slackutil.GetUsername(apiClient, groupId)
				cachedUsernames[groupId] = username
				groupNames = append(groupNames, username)
			} else {
				groupNames = append(groupNames, username)
			}
		}

		body = append(body, scheduleBody{
			Id:         session.Id,
			Username:   cachedUsernames[session.UserId],
			GroupNames: groupNames,
			Machine:    session.Machine,
			Reason:     session.Reason,
			Duration:   int(session.Duration),
			Time:       session.Time.Format(time.RFC3339),
		})
	}

	// generate body and 500 on fail
	j, err := json.Marshal(body)
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
