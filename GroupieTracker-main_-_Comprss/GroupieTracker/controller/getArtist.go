package controller

import (
	"GroupieTracker/model"
	"encoding/json"
	"fmt"
)

func GetArtistsApi() (*[]model.ArtistsApi, error) {

	body, err := MakeRequest(`https://groupietrackers.herokuapp.com/api/artists`)
	if err != nil {
		return nil, err
	}
	s, err := ParseArtistsApi(body)
	return s, err
}

func ParseArtistsApi(body []byte) (*[]model.ArtistsApi, error) {
	var artist = new([]model.ArtistsApi)
	err := json.Unmarshal(body, artist)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return artist, err
}
