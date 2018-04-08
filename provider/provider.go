package provider

import (
	"fmt"
	//	"fresh_cve/cve"
	//	"fresh_cve/cve/cve_redhat"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Provider struct {
	Address    string
	TimeoutSec int
	Proxy      string
	Format     string
}

type Answer struct {
	Format string
	Body   string
}

func (p Provider) GetBodyToChan(c chan Answer) error {
	b, err := p.GetBody()
	if err != nil {
		return err
	}
	c <- Answer{Format: p.Format, Body: b}
	return nil
}

func (p Provider) GetBody() (string, error) {
	t1 := time.Now()
	client := &http.Client{}
	var proxyUrl *url.URL
	var err error
	if p.Proxy != "" {
		proxyUrl, err = url.Parse(p.Proxy)
		if err != nil {
			return "", err
		}
		client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	} else {
		log.Println("Do you use Proxy server? check config.json -> Proxy settings")
	}
	resp, err := client.Get(p.Address)
	log.Println("Get", p.Address)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	log.Println("resp.StatusCode", resp.StatusCode)
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("resp.StatusCode=%d", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	log.Println("response time=", time.Now().Sub(t1))
	return string(body), err
}

//var Proxy string = "http://omin-e:Bercut19@192.168.55.8:3128/"

//func GetCves(Providers []Provider) (results []string) {
//	c := make(chan string, len(Providers))
//	for _, p := range Providers {
//		go GetToChan(p.Address, c)
//	}
//	timeout := time.After(80 * time.Second)
//	for _, _ = range Providers {
//		select {
//		case result := <-c:
//			results = append(results, result)
//		case <-timeout:
//			log.Println("timed out")
//			return results
//		}
//	}
//	return results
//}
//func GetToChan(address string, c chan string) {
//	if ans, err := Get(address); err == nil {
//		c <- ans
//	}
//}

//func (p Provider) GetCves() (list []cve.Cve, err error) {
//	body, err := GetBody(p.Address)
//	if err != nil {
//		return list, err
//	}
//	if p.Format == "redhat" {
//		list, err := cve_redhat.Decode(body)
//		return list, err
//	}
//	//	if p.Format == "circl"
//	return list, nil
//}

//func (p Provider) GetBodyTimeout() (string, error) {
//	//решил делать раздельные timeout (ТЗ если провайдер не ответил в течении 2 х секунд. *Вынести таймауты в conf файл
//	timeout := time.After(time.Duration(p.TimeoutSec) * time.Second)
//	c := make(chan Answer)
//	go p.GetBodyToChan(c)
//	select {
//	case b := <-c:
//		return b, nil
//	case <-timeout:
//		return "", fmt.Errorf("Timout get responce from %s", p.Address)
//	}
//}

//func GetBody(address string) (string, error) {
//	client := &http.Client{}
//	var proxyUrl *url.URL
//	var err error
//	if Proxy != "" {
//		proxyUrl, err = url.Parse(Proxy)
//		if err != nil {
//			return "", err
//		}
//		client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
//	}
//	resp, err := client.Get(address)
//	if err != nil {
//		return "", err
//	}
//	defer resp.Body.Close()
//	log.Println("resp.StatusCode", resp.StatusCode)
//	if resp.StatusCode != 200 {
//		return "", fmt.Errorf("resp.StatusCode=%d", resp.StatusCode)
//	}
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return "", err
//	}
//	return string(body), err
//}
