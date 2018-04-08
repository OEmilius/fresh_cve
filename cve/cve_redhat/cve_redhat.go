package cve_redhat

import (
	"encoding/json"
	"fresh_cve/cve"
	"log"
)

//type Cve struct {
//	Id        string `json:"CVE"`
//	Published string `json:"public_date"`
//	Summary   string `json:"bugzilla_description"`
//	Url       string `json:"resource_url"`
//}

//func (c Cve) String() {
//	fmt.Sprintf("%v", c)
//}

//func Decode(s string) (c Cve) {
//	if err := json.Unmarshal([]byte(one_cve), &c); err != nil {
//		panic(err)
//	}
//	return c
//}

func Decode(s string) (cve_list []cve.Cve, err error) {
	var list []map[string]interface{}
	if err := json.Unmarshal([]byte(s), &list); err != nil {
		log.Println("decode error", err)
		return cve_list, err
	}
	for _, m := range list {
		cve := decode_one(m)
		cve_list = append(cve_list, cve)
	}
	return cve_list, nil
}

func decode_one(m map[string]interface{}) (c cve.Cve) {
	c.Id = m["CVE"].(string)
	c.Published = m["public_date"].(string)
	c.Summary = m["bugzilla_description"].(string)
	c.Urls = append(c.Urls, m["resource_url"].(string))
	return c
}
