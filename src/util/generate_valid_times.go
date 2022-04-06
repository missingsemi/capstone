package util

import (
	"sync"
	"time"

	"github.com/missingsemi/capstone/database"
	"github.com/missingsemi/capstone/model"
)

func GenerateValidTimes(machine model.Machine, duration int) []time.Time {
	now := time.Unix((time.Now().Unix()/(15*60)+1)*15*60, 0)
	lowerBound := now.Add(-4 * time.Hour)
	sessions, _ := database.GetSessionsAfterTimeByMachine(lowerBound, machine.Id)

	var numGroups int
	if len(sessions) != 0 {
		numGroups = len(sessions)/10 + 1
	}
	wg := sync.WaitGroup{}
	wg.Add(numGroups)
	c := make(chan []int8, numGroups)

	for i := 0; i < len(sessions); i += 10 {
		if (i + 10) > len(sessions) {
			go applySessions(sessions[i:], lowerBound, now, c, &wg)
		} else {
			go applySessions(sessions[i:i+10], lowerBound, now, c, &wg)
		}
	}

	wg.Wait()

	timeSlots := make([]int8, 96)
	for i := 0; i < numGroups; i++ {
		partial := <-c
		for j, v := range partial {
			timeSlots[j] += v
		}
	}

	availableTimes := make([]time.Time, 0)
	count := 0
	for i, v := range timeSlots {
		if v < int8(machine.Count) {
			count++
		} else {
			count = 0
		}
		if count >= duration/15 {
			offset := 15 * (i - duration/15) * int(time.Minute)
			t := now.Add(time.Duration(offset))
			availableTimes = append(availableTimes, t)
		}
		if len(availableTimes) == 10 {
			break
		}
	}

	return availableTimes
}

func applySessions(sessions []model.ScheduleSession, lowerBound time.Time, now time.Time, c chan []int8, wg *sync.WaitGroup) {
	// 24 hours worth of time slots
	timeSlots := make([]int8, 96)

	for _, session := range sessions {
		lowerIndex := toIndex(now, session.Time)
		blocks := session.Duration / 15
		for i := lowerIndex; i < lowerIndex+blocks; i++ {
			if i < 0 {
				continue
			}
			timeSlots[i]++
		}
	}

	c <- timeSlots
	wg.Done()
}

func toIndex(now time.Time, t time.Time) int64 {
	// Difference between t and now
	// Counted in 15 minute segments
	return (t.Unix() - now.Unix()) / (15 * 60)
}
