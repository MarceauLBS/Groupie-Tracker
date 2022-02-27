package controller

import (
	"GroupieTracker/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var artist *[]model.ArtistsApi
var dates *model.Dates
var relation *model.Relation

var colorGreen = "\033[32m"
var colorReset = "\033[0m"

func MakeRequest(url string) ([]byte, error) {

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return []byte(body), err
}

func artistsInit() {
	artistsData, err := GetArtistsApi()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(colorGreen), "Artists done.")
	fmt.Print(string(colorReset))
	artist = artistsData
}
func datesInit() {
	datesData, err := GetConcertDates()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(colorGreen), "Dates done.")
	fmt.Print(string(colorReset))
	dates = datesData
}
func relationInit() {
	relationData, err := GetRelation()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(colorGreen), "Relation done.")
	fmt.Print(string(colorReset))
	relation = relationData
}

func GetDataByID(id int, shown bool) *model.Artist {

	var artistData = new(model.Artist)
	artistData.Id = int64(id)
	artistData.Name = (*artist)[id].Name
	artistData.Image = (*artist)[id].Image

	for i, s := range dates.Index[id].Dates {
		if string(dates.Index[id].Dates[i][0]) == "*" {
			dates.Index[id].Dates[i] = s[1:]
		}
	}

	artistData.ConcertDates = dates.Index[id].Dates
	var countries []string
	var cities []string
	for i, s := range countries {
		countries[i] = strings.TrimPrefix(s, "-")
	}
	for i, s := range cities {
		cities[i] = strings.TrimSuffix(s, "-")
	}

	empJSON, err := json.MarshalIndent(artistData, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	if shown {
		fmt.Printf(string(empJSON))
	}
	return artistData

}

func Init() {

	start := time.Now()
	fmt.Printf("\n")

	go artistsInit()
	go datesInit()
	go relationInit()

	time.Sleep(600 * time.Millisecond)
	elapsed := time.Since(start)
	fmt.Printf("\ntook %s to gather data\n", elapsed)

}
