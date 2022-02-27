package controller

import (
	"GroupieTracker/model"
	"encoding/json"
	"fmt"
)

func GetConcertDates() (*model.Dates, error) {

	body, err := MakeRequest("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		return nil, err
	}
	s, err := ParseConcertDates(body)
	return s, err
}

func ParseConcertDates(body []byte) (*model.Dates, error) {
	var dates = new(model.Dates)
	err := json.Unmarshal(body, dates)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return dates, err
}
