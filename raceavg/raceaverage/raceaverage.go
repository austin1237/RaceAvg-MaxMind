package raceaverage

import (
	"errors"
	"math"
	"strconv"
	"strings"
	"time"
)

type raceTime struct {
	hours   int
	minutes int
	ampm    string
	day     int
}

func AvgMinutes(rawRaceStrings []string) (int, error) {
	var durations []float64
	for _, raceString := range rawRaceStrings {
		rt, err := convertStrToRaceTime(raceString)
		if err != nil {
			return 0, err
		}
		duration := getRaceDuration(rt)
		durations = append(durations, duration)

	}

	average := averageDurations(durations)
	return int(math.Round(average)), nil
}

func averageDurations(durations []float64) float64 {
	total := 0.0
	for _, duration := range durations {
		total += duration
	}
	return total / float64(len(durations))
}

func convertStrToRaceTime(rawRaceTime string) (raceTime, error) {
	rt := raceTime{}
	splitStrings := strings.Split(rawRaceTime, " ")
	hoursAndMinutes := splitStrings[0]
	splitHoursAndMinutes := strings.Split(hoursAndMinutes, ":")
	rt.ampm = splitStrings[1]
	day, err := strconv.Atoi(splitStrings[3])

	if err != nil {
		err = errors.New("Unable to parse day in date string " + rawRaceTime)
		return rt, err
	}
	rt.day = day

	hours, err := strconv.Atoi(splitHoursAndMinutes[0])
	if err != nil {
		err = errors.New("Unable to parse hour in date string " + rawRaceTime)
		return rt, err
	}

	minutes, err := strconv.Atoi(splitHoursAndMinutes[1])
	if err != nil {
		err = errors.New("Unable to parse minutes in date string " + rawRaceTime)
		return rt, err
	}
	rt.hours = hours
	rt.minutes = minutes
	return rt, nil
}

func getRaceDuration(rt raceTime) float64 {
	if rt.ampm == "AM," && rt.hours == 12 {
		rt.hours = 0
	} else if rt.ampm == "PM," && rt.hours != 12 {
		rt.hours = 12 + rt.hours
	}

	//Calculates the minutes between two date where the start date is JAN 1 At 8am(start of the race)
	// And the end time which is JAN rt.day AT rt.hours, rt.minutes (when they finished)
	startTime := time.Date(1970, time.January, 1, 8, 0, 0, 0, time.UTC)
	endTime := time.Date(1970, time.January, rt.day, rt.hours, rt.minutes, 0, 0, time.UTC)
	duration := endTime.Sub(startTime).Minutes()
	return duration
}
