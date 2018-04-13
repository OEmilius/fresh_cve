//Convert json array to array []cve{} format
package cve_redhat

import (
	"encoding/json"
	"log"

	"github.com/OEmilius/fresh_cve/cve"
)

//Decode from string containing json array to []cve.Cve
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
