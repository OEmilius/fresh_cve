//html  client, send get to address
package provider

import (
	"fmt"
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
	log.Println("start ProviderGetBodyToChan")
	b, err := p.GetBody()
	if err != nil {
		log.Println("ProviderGetBodyToChan", err)
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
			log.Println(`example in config.json "Proxy": "http://user:pass@10.10.10.10:3128/",`)
			log.Fatalln(err)
			return "", err
		}
		client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
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
