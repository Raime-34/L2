package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func events_for_day(w http.ResponseWriter, req *http.Request) {
	resp := GetResult{
		Result: make([]Event, 0),
	}
	// result := make([]Event, 0)
	date := time.Now().Format(dateFormat)

	for _, event := range events {
		if event.Date == date {
			resp.Result = append(resp.Result, event)
		}
	}

	fmt.Println(resp)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func events_for_week(w http.ResponseWriter, req *http.Request) {
	result := make([]Event, 0)
	leftBorder := time.Now().AddDate(0, 0, -1)
	rightBorder := leftBorder.AddDate(0, 0, 8)

	for _, event := range events {
		date, err := time.Parse(dateFormat, event.Date)
		if err != nil {
			continue
		}

		if date.After(leftBorder) && rightBorder.After(date) {
			result = append(result, event)
		}
	}

	json.NewEncoder(w).Encode(GetResult{
		Result: result,
	},
	)
}

func events_for_month(w http.ResponseWriter, req *http.Request) {
	result := make([]Event, 0)
	leftBorder := time.Now().AddDate(0, 0, -1)
	rightBorder := leftBorder.AddDate(0, 1, 0)

	for _, event := range events {
		date, err := time.Parse(dateFormat, event.Date)
		if err != nil {
			continue
		}

		if date.After(leftBorder) && rightBorder.After(date) {
			result = append(result, event)
		}
	}

	json.NewEncoder(w).Encode(GetResult{
		Result: result,
	},
	)
}

func create_event(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ReqError{
			ReqError: "Ошибка",
		},
		)
		return
	}

	if !req.Form.Has("name") || !req.Form.Has("date") {
		json.NewEncoder(w).Encode(ReqError{
			ReqError: "Введены невалидные данные",
		},
		)
		return
	}

	name := req.Form.Get("name")
	date := req.Form.Get("date")

	events[int(time.Now().Unix())] = Event{
		Name: name,
		Date: date,
	}

	json.NewEncoder(w).Encode(PostResult{
		Result: "Жест",
	},
	)
}

func update_event(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		return
	}

	name := req.Form.Get("name")
	date := req.Form.Get("date")

	if name == "" || date == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ReqError{
			ReqError: "Введены невалидные данные",
		},
		)
		return
	}

	for _, event := range events {
		if event.Name == name {
			event.Date = date
		}
	}

	json.NewEncoder(w).Encode(PostResult{
		Result: "Жест",
	},
	)
}
