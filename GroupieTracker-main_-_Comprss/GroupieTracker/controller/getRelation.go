package controller

import (
	"GroupieTracker/model"
	"encoding/json"
	"fmt"
)

func GetRelation() (*model.Relation, error) {

	body, err := MakeRequest("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return nil, err
	}
	s, err := ParseRelation(body)
	return s, err
}

func ParseRelation(body []byte) (*model.Relation, error) {
	var relation = new(model.Relation)
	err := json.Unmarshal(body, relation)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return relation, err
}
