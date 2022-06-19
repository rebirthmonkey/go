package main

import (
	"fmt"
	"time"
)

func timeUntilSchedule(schedule string) (time.Duration, error) {
	now := time.Now().UTC()
	layout := "2006-01-02T15:04:05Z"
	s, err := time.Parse(layout, schedule)
	if err != nil {
		return time.Duration(0), err
	}
	return s.Sub(now), nil
}

func main() {
	d, _ := timeUntilSchedule("2022-06-12T02:58:31Z")
	fmt.Println("the duration is", d)
}
