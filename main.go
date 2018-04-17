//how it work
//- read config.json
//- init cache
//- load old cache
//- init LOADER
//- start web server
//- periodically check fresh cve

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/OEmilius/fresh_cve/cache"
	"github.com/OEmilius/fresh_cve/grpc_server"
	"github.com/OEmilius/fresh_cve/loader"
	"github.com/OEmilius/fresh_cve/webserver"
)

type Config struct {
	ServAddr          string       // address for start web server
	GrpcPort          string       // Port for listen grps server
	Interval          int          // sec for PeriodicalyLoad
	DefaultTimeoutSec int          // sec for Loader
	Proxy             string       // for providers proxy, if no proxy use ""
	FileGobName       string       //cache // for gob file // where save gob file
	Sources           []loader.Src //sources {"Format": "redh","Address":"http://ya.ru"}
}

var cfg Config

func main() {
	log.Println("start")
	runtime.GOMAXPROCS(runtime.NumCPU())
	cfg = readConfig("config.json")
	log.Println("config=", cfg)
	CACHE := cache.NewCache()
	CACHE.FileGobName = cfg.FileGobName
	err := CACHE.Load()
	if err != nil {
		log.Println("error loading cache:", err)
	}
	defer CACHE.Save()
	log.Println("cache from file len=", len(CACHE.GetAllCve()))
	LOADER := loader.NewLoader()
	LOADER.DefaultTimeoutSec = cfg.DefaultTimeoutSec
	LOADER.Sources = cfg.Sources
	LOADER.Proxy = cfg.Proxy
	webserver.ServAddr = cfg.ServAddr
	webserver.CACHE = CACHE
	go webserver.Start()
	grpc_server.ListenPort = cfg.GrpcPort
	grpc_server.CACHE = CACHE
	go grpc_server.Start()
	go PeriodicalyLoad(LOADER, CACHE)
	fmt.Scanln()
}

func PeriodicalyLoad(l *loader.Loader, cache *cache.Cache) {
	Tick := time.NewTicker(time.Duration(cfg.Interval) * time.Second)
	for {
		<-Tick.C
		log.Println("time to get fresh CVE| cfg.Interval=", cfg.Interval)
		cache.AddList(l.Query())
		log.Println("cache size is=", len(cache.GetAllCve()))
	}
}

func readConfig(fname string) Config {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatalln("opening config file error", err)
	}
	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalln("error reading config file", err)
	}
	return config
}
