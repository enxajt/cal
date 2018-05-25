package main

import (
	"fmt"
	"github.com/enxajt/cal/Utils"
)

func main() {
	datas := Utils.Read("schedule.csv")
	calendarId := Utils.GetCalendarId()
	for _, schedule := range datas {
		s := Utils.SetSchedule(schedule)
		Utils.CreateEvent(s, calendarId)
	}

	fmt.Println("finished!")
}
