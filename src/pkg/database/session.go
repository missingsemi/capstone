package database

import (
	"strings"
	"time"

	"github.com/missingsemi/capstone/pkg/model"
)

func GetSessions() ([]model.ScheduleSession, error) {
	rows, err := db.Query("SELECT * FROM schedule;")
	if err != nil {
		return []model.ScheduleSession{}, err
	}
	defer rows.Close()

	sessions := make([]model.ScheduleSession, 0)

	for rows.Next() {
		session := model.ScheduleSession{}
		var datetime string
		var groupIds string
		err := rows.Scan(&session.Id, &session.UserId, &groupIds, &session.Machine, &session.Reason, &session.Duration, &datetime, &session.Stage)
		if err != nil {
			return sessions, err
		}
		dateObj, err := time.Parse(time.RFC3339, datetime)
		if err != nil {
			return sessions, err
		}
		session.Time = dateObj
		splitIds := strings.Split(groupIds, ",")
		session.GroupIds = splitIds
		sessions = append(sessions, session)
	}

	return sessions, nil
}

func GetSessionsAfterTimeByMachine(datetime time.Time, machineId string) ([]model.ScheduleSession, error) {
	stmt, err := db.Prepare("SELECT * FROM schedule WHERE time > ? AND machine_id = ?;")
	if err != nil {
		return []model.ScheduleSession{}, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(datetime.Format(time.RFC3339), machineId)
	if err != nil {
		return []model.ScheduleSession{}, err
	}
	defer rows.Close()

	sessions := make([]model.ScheduleSession, 0)

	for rows.Next() {
		session := model.ScheduleSession{}
		var datetime string
		var groupIds string
		err := rows.Scan(&session.Id, &session.UserId, &groupIds, &session.Machine, &session.Reason, &session.Duration, &datetime, &session.Stage)
		if err != nil {
			return sessions, err
		}
		dateObj, err := time.Parse(time.RFC3339, datetime)
		if err != nil {
			return sessions, err
		}
		session.Time = dateObj
		splitIds := strings.Split(groupIds, ",")
		session.GroupIds = splitIds
		sessions = append(sessions, session)
	}

	return sessions, nil
}

func GetSessionById(id int) (model.ScheduleSession, error) {
	stmt, err := db.Prepare("SELECT * FROM schedule WHERE id = ?;")
	if err != nil {
		return model.ScheduleSession{}, err
	}
	defer stmt.Close()

	session := model.ScheduleSession{}
	var datetime string
	var groupIds string
	err = stmt.QueryRow(id).Scan(&session.Id, &session.UserId, &groupIds, &session.Machine, &session.Reason, &session.Duration, &datetime, &session.Stage)
	if err != nil {
		return model.ScheduleSession{}, err
	}
	dateObj, err := time.Parse(time.RFC3339, datetime)
	if err != nil {
		return model.ScheduleSession{}, err
	}
	session.Time = dateObj
	splitIds := strings.Split(groupIds, ",")
	session.GroupIds = splitIds
	return session, nil
}

func CreateSession(session model.ScheduleSession) error {
	stmt, err := db.Prepare("INSERT INTO schedule (user_id, group_ids, machine_id, reason, duration, time) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	groupIds := ""
	for _, id := range session.GroupIds {
		groupIds += id
	}

	_, err = stmt.Exec(session.UserId, groupIds, session.Machine, session.Reason, session.Duration, session.Time.Format(time.RFC3339))
	return err
}

func ModifySession(id int, session model.ScheduleSession) error {
	stmt, err := db.Prepare("UPDATE schedule SET user_id = ?, group_ids = ?, machine_id = ?, reason = ?, duration = ?, time = ?, stage = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	groupIds := ""
	for _, id := range session.GroupIds {
		groupIds += id
	}

	_, err = stmt.Exec(session.UserId, groupIds, session.Machine, session.Reason, session.Duration, session.Time.Format(time.RFC3339), session.Stage, id)
	return err
}

func DeleteSession(id int) error {
	stmt, err := db.Prepare("DELETE FROM schedule WHERE id = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}

func GetUnfinishedSessions() ([]model.ScheduleSession, error) {
	rows, err := db.Query("SELECT * FROM schedule WHERE NOT stage = 4;")
	if err != nil {
		return []model.ScheduleSession{}, err
	}
	defer rows.Close()

	sessions := make([]model.ScheduleSession, 0)

	for rows.Next() {
		session := model.ScheduleSession{}
		var datetime string
		var groupIds string
		err := rows.Scan(&session.Id, &session.UserId, &groupIds, &session.Machine, &session.Reason, &session.Duration, &datetime, &session.Stage)
		if err != nil {
			return sessions, err
		}
		dateObj, err := time.Parse(time.RFC3339, datetime)
		if err != nil {
			return sessions, err
		}
		session.Time = dateObj
		splitIds := strings.Split(groupIds, ",")
		session.GroupIds = splitIds
		sessions = append(sessions, session)
	}

	return sessions, nil
}

func GetUpcomingSessionsByUser(userId string) ([]model.ScheduleSession, error) {
	stmt, err := db.Prepare("SELECT * FROM schedule WHERE time > ? AND user_id = ? ORDER BY time ASC;")
	if err != nil {
		return []model.ScheduleSession{}, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(time.Now().Format(time.RFC3339), userId)
	if err != nil {
		return []model.ScheduleSession{}, err
	}
	defer rows.Close()

	sessions := make([]model.ScheduleSession, 0)

	for rows.Next() {
		session := model.ScheduleSession{}
		var datetime string
		var groupIds string
		err := rows.Scan(&session.Id, &session.UserId, &groupIds, &session.Machine, &session.Reason, &session.Duration, &datetime, &session.Stage)
		if err != nil {
			return sessions, err
		}
		dateObj, err := time.Parse(time.RFC3339, datetime)
		if err != nil {
			return sessions, err
		}
		session.Time = dateObj
		splitIds := strings.Split(groupIds, ",")
		session.GroupIds = splitIds
		sessions = append(sessions, session)
	}

	return sessions, nil
}

func GetSessionsBetweenTimes(start time.Time, end time.Time) ([]model.ScheduleSession, error) {
	stmt, err := db.Prepare("SELECT * FROM schedule WHERE time > ? AND time < ? ORDER BY time ASC;")
	if err != nil {
		return []model.ScheduleSession{}, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(start.Format(time.RFC3339), end.Format(time.RFC3339))
	if err != nil {
		return []model.ScheduleSession{}, err
	}
	defer rows.Close()

	sessions := make([]model.ScheduleSession, 0)

	for rows.Next() {
		session := model.ScheduleSession{}
		var datetime string
		var groupIds string
		err := rows.Scan(&session.Id, &session.UserId, &groupIds, &session.Machine, &session.Reason, &session.Duration, &datetime, &session.Stage)
		if err != nil {
			return sessions, err
		}
		dateObj, err := time.Parse(time.RFC3339, datetime)
		if err != nil {
			return sessions, err
		}
		session.Time = dateObj
		splitIds := strings.Split(groupIds, ",")
		session.GroupIds = splitIds
		sessions = append(sessions, session)
	}

	return sessions, nil
}
