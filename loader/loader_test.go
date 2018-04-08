package loader

import (
	"fmt"
)

func ExampleLoader_QueryAndCombine() {
	l := NewLoader()
	l.DefaultTimeoutSec = 2
	l.Proxy = ""
	src1 := Src{Address: "http://cve.circl.lu/api/last/5", Format: "circl"}
	src2 := Src{Format: "redh", Address: "http://access.redhat.com/labs/securitydataapi/cve.json?after=2018-03-14"}
	l.Sources = []Src{src2, src1}
	result := l.QueryAndCombine()
	fmt.Println(len(result))
	for i, c := range result {
		fmt.Println(i, c.Id, c.Published)
	}
	//Output: 2
}

//func ExampleLoader_Query() {
//	l := Loader{}
//	l.DefaultTimeoutSec = 2
//	l.Proxy = ""
//	src1 := Src{Address: "http://cve.circl.lu/api/last/5", Format: "circl"}
//	src2 := Src{Format: "redh", Address: "http://access.redhat.com/labs/securitydataapi/cve.json?after=2018-03-14"}
//	l.Sources = []Src{src2, src1}
//	result := l.Query()
//	fmt.Println(len(result))
//	//Output: 2
//}
