package model

type Artist struct {
	Id           int64
	Name         string
	Image        string
	Members      []string
	CreationDate uint16
	ConcertDates []string
	Relation     map[string][]string
	Countries    []string
	Cities       []string
}
