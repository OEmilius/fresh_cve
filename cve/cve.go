//common cve format

package cve

type Cve struct {
	Id        string `json:"id"`
	Published string `json:"published"`
	Summary   string `json:"summary"`
	Urls      []string
}
