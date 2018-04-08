/*
[
    {
        "Modified": "2018-04-04T21:29:07.283000",
        "Published": "2018-04-04T21:29:07.267000",
        "cvss": null,
        "cwe": "Unknown",
        "id": "CVE-2018-9309",
        "last-modified": "2018-04-04T21:29:07.283000",
        "references": [
            "https://github.com/lihonghuyang/vulnerability/blob/master/dl_sendsms.php.md"
        ],
        "summary": "An issue was discovered in zzcms 8.2. It allows SQL injection via the id parameter in a dl/dl_sendsms.php request.",
        "vulnerable_configuration": [],
        "vulnerable_configuration_cpe_2_2": []
    },

...
}
]

http://cve.circl.lu/api/last/14 вернёт последние 3
*/

package cve_circl

import (
	"encoding/json"
	//"fmt"
	//"time"
	"fresh_cve/cve"
	"log"
)

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
	c.Id = m["id"].(string)
	c.Published = m["Published"].(string)
	c.Summary = m["summary"].(string)
	//c.Urls = append(c.Urls, m["resource_url"].(string))
	return c
}

//func Decode(s string) (c Cve, err error) {
//	if err = json.Unmarshal([]byte(s), &c); err != nil {
//		return c, err
//	}
//	return c, nil
//}

//func DecodeArray(s string) (c_array []Cve, err error) {
//	if err = json.Unmarshal([]byte(s), &c_array); err != nil {
//		return c_array, err
//	}
//	return c_array, nil
//}

//func (c Cve) String() string {
//	return fmt.Sprintf("%#v", c)
//}
