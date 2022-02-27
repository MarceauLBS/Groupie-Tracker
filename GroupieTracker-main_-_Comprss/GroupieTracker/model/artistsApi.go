package model

type ArtistsApi struct {
	Id           int64    `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate uint16   `json:"creationDate"`
}
