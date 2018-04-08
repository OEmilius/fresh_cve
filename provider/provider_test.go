package provider

import (
	"fmt"
	"testing"
	//	"time"
)

//func TestGet(t *testing.T) {
//	t1 := time.Now()
//	url := "http://access.redhat.com/labs/securitydataapi/cve.json?after=2018-03-14"
//	answer, err := Get(url)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println("answer len=", len(answer))
//	t2 := time.Now()
//	fmt.Println("execution time", t2.Sub(t1))
//}

//func TestGetToChan(t *testing.T) {
//	address := "http://access.redhat.com/labs/securitydataapi/cve.json?after=2018-03-14"
//	c := make(chan string, 2)
//	GetToChan(address, c)
//	fmt.Println(<-c)
//}
//func ExampleProvider_GetBodyTimeout() {
//	p := Provider{}
//	p.Address = "http://access.redhat.com/labs/securitydataapi/cve.json?after=2018-03-14"
//	p.TimeoutSec = 20
//	b, err := p.GetBodyTimeout()
//	if b == "" {
//		fmt.Println("empty body")
//	}
//	fmt.Println(err)
//	//Output: <nil>
//}

//func ExampleProvicer_GetBody() {
//	p := Provider{}
//	p.Address = "http://access.redhat.com/labs/securitydataapi/cve.json?after=2018-03-14"
//	b, err := p.GetBody()
//	if b == "" {
//		fmt.Println("empty body")
//	}
//	fmt.Println(err)
//	//Output: <nil>
//}
func TestProvider(t *testing.T) {
	p := Provider{}
	p.Address = "http://access.redhat.com/labs/securitydataapi/cve.json?after=2018-03-14"
	fmt.Println(p.Address)
}

//func TestGetSves(t *testing.T) {
//	p := Provider{}
//	p.Address = "http://access.redhat.com/labs/securitydataapi/cve.json?after=2018-03-14"
//	p.Format = "redhat"
//	list, err := p.GetCves()
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(list)
//	}

//}
