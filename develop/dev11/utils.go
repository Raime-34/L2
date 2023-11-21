package main

import (
	"time"
)

const (
	dateFormat = "2006-01-02"
)

func initEvents() {

	events[0] = Event{
		Name: "Тест 1",
		Date: time.Now().Format(dateFormat),
	}

	events[1] = Event{
		Name: "Тест 2",
		Date: time.Now().AddDate(0, 1, 1).Format(dateFormat),
	}

	events[2] = Event{
		Name: "Тест 3",
		Date: time.Now().AddDate(0, 2, 21).Format(dateFormat),
	}

	events[3] = Event{
		Name: "Тест 4",
		Date: time.Now().AddDate(0, 3, 31).Format(dateFormat),
	}

	events[4] = Event{
		Name: "Тест 5",
		Date: time.Now().AddDate(0, 0, 5).Format(dateFormat),
	}

	events[5] = Event{
		Name: "Тест 6",
		Date: time.Now().AddDate(0, 0, 21).Format(dateFormat),
	}
}
