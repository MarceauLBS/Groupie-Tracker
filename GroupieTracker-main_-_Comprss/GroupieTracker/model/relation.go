package model

type Relation struct {
	Index []struct {
		Id             uint64              `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}
