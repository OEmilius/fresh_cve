package cve

//_ "encoding/json"

//один на всех мы за ценой не постоим

type Cve struct {
	Id        string `json:"id"`
	Published string `json:"published"`
	Summary   string `json:"summary"`
	Urls      []string
}

// выше формат маршалинга ответных
