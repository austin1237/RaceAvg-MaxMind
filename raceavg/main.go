package main

import (
	"log"
	"os"

	"github.com/user/raceavg/raceaverage"
)

func main() {
	log.Println("Race Averages Started")
	raceStrings := []string{"02:00 PM, DAY 19",
		"02:00 PM, DAY 20",
		"01:58 PM, DAY 20",
	}
	average, err := raceaverage.AvgMinutes(raceStrings)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	log.Printf("Average is %v exiting", average)
	os.Exit(0)
}
