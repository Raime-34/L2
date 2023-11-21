package main

type Event struct {
	Name string
	Date string
}

type GetResult struct {
	Result []Event
}

type PostResult struct {
	Result string
}

type ReqError struct {
	ReqError string
}
