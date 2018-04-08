package loader

import (
	"fresh_cve/cache"
	"fresh_cve/cve"
	circl "fresh_cve/cve/cve_circl"
	redhat "fresh_cve/cve/cve_redhat"
	"fresh_cve/provider"
	"log"
	"time"
)

type Loader struct {
	Sources           []Src
	DefaultTimeoutSec int
	Proxy             string
	Cache             *cache.Cache
}

type Src struct {
	Format  string
	Address string
}

func DecodeAnswer(a provider.Answer) (result []cve.Cve) {
	var err error
	switch a.Format {
	case "redh":
		log.Println("redhat format")
		//log.Println("a.Body=", a.Body)
		result, err = redhat.Decode(a.Body)
		if err != nil {
			log.Println(err)
			return result
		}
	case "circl":
		log.Println("circl Format")
		result, err = circl.Decode(a.Body)
		if err != nil {
			log.Println(err)
		}
		return result
	}
	return result
}

func NewLoader() *Loader {
	return &Loader{Cache: cache.NewCache()}
}

func (l *Loader) QueryAndCombine() []cve.Cve {
	l.Cache.AddList(l.Query())
	return l.Cache.GetAllCve()
}

//уточнить по тз реализовать default timeout или для каждого свой
func (l Loader) Query() (result []cve.Cve) {
	log.Println("Loader Query start")
	ans_chan := make(chan provider.Answer, len(l.Sources))
	timeout := time.After(time.Duration(l.DefaultTimeoutSec) * time.Second)
	log.Println("loader timeout=", l.DefaultTimeoutSec)
	for _, src := range l.Sources {
		p := provider.Provider{}
		p.Address = src.Address
		p.Format = src.Format
		p.Proxy = l.Proxy
		go p.GetBodyToChan(ans_chan)
	}
	for i, _ := range l.Sources {
		log.Println(i)
		select {
		case a := <-ans_chan:
			log.Println("get answer Format = ", a.Format)
			answers := DecodeAnswer(a)
			log.Println("answers count=", len(answers))
			result = append(result, answers...)
		case <-timeout:
			log.Println("Default timeout ")
		}
	}
	return result
}
