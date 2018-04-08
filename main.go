package main

import (
	"encoding/json"
	"fmt"
	"fresh_cve/cache"
	"fresh_cve/loader"
	"fresh_cve/webserver"
	"log"
	"os"
	"runtime"
	"time"
)

type Config struct {
	ServAddr          string       //:8081
	Interval          int          // как часто загружать новые уязвимости
	DefaultTimeoutSec int          //in sec
	Proxy             string       // for providers proxy
	FileGobName       string       //cache // for gob file // where save gob file
	Sources           []loader.Src //источники уязвимостей {"Format": "redh","Address":"http://ya.ru"}
}

var cfg Config

/*
- читаем конфигурационный фаил
- инициализируем кэш
- загружем ранее сохраненный кэш
- инициализирем LOADER - (опрашивает провайдеров)
- передаем в web сервер адрес инициализированного кэша
- запускаем web сервер
- запускаем пириодическу загрузку свежих уязвимостей
*/
func main() {
	fmt.Println("start")
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(0)
	log.SetOutput(new(logWriter))
	cfg = ReadConfig("config.json")
	log.Println("config=", cfg)
	CACHE := cache.NewCache()
	CACHE.FileGobName = cfg.FileGobName
	err := CACHE.Load()
	if err != nil {
		log.Println("error loading cache:", err)
	}
	defer CACHE.Save()
	log.Println("cache from file len=", len(CACHE.GetAllCve()))
	//log.Println(CACHE.GetAllcveJson())
	LOADER := loader.NewLoader()
	LOADER.DefaultTimeoutSec = cfg.DefaultTimeoutSec
	LOADER.Sources = cfg.Sources
	LOADER.Proxy = cfg.Proxy
	//CACHE.AddList(LOADER.Query())
	webserver.ServAddr = cfg.ServAddr
	webserver.CACHE = CACHE //передаем созданный кэш что бы веб сервер мог дергать метод
	// GetAllcveJson()
	go webserver.Start_and_open()
	//go webserver.Start()
	go PeriodicalyLoad(LOADER, CACHE)
	fmt.Scanln()
}

func PeriodicalyLoad(l *loader.Loader, cache *cache.Cache) {
	Tick := time.NewTicker(time.Duration(cfg.Interval) * time.Second)
	for {
		<-Tick.C
		log.Println("time to get fresh CVE")
		cache.AddList(l.Query())
		log.Println("cache size is=", len(cache.GetAllCve()))
	}
}

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().UTC().Format("15:04:05.999") + " " + string(bytes))
}

func ReadConfig(fname string) Config {
	runtime.GOMAXPROCS(runtime.NumCPU())
	file, err := os.Open(fname)
	if err != nil {
		//dbg.Println("opening config file error", err)
		log.Fatalln("opening config file error", err)
		panic(err)
	}
	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		//dbg.Println("error reading config file", err)
		log.Fatalln("error reading config file", err)
		panic(err)
	}
	//fmt.Println("config=", config)
	return config
}
