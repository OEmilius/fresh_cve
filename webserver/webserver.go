package webserver

import (
	"fresh_cve/cache"
	"log"
	"net/http"
	"os/exec"
)

var CACHE *cache.Cache
var ServAddr string = ":8081" // listen & serv address

func serveCves(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api/v1/cves" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	log.Println("GET /api/v1/cves from:", r.Host)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// TODO: в случае ошибки на нижнем уровне выводить сюда ошибку
	w.Write([]byte(CACHE.GetAllcveJson()))
}

func Start() {
	log.Println("http.ListenAndServe:", ServAddr)
	http.HandleFunc("/api/v1/cves", serveCves)
	if err := http.ListenAndServe(ServAddr, nil); err != nil {
		log.Println(err)
	}
}

func Start_and_open() {
	go Start()
	_ = exec.Command("cmd", "/c", "start", "http://localhost:8081/api/v1/cves").Start()
}
