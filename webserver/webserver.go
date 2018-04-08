package webserver

import (
	"fresh_cve/cache"
	"log"
	"net/http"
	"os/exec"
)

var CACHE *cache.Cache
var ServAddr string = ":8081" // адрес на котором запустится web server

//func main() {
//	go Start_and_open()
//	fmt.Scanln()
//}

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

	//	var v = struct {
	//		Host string
	//		Data string
	//	}{
	//		r.Host,
	//		"some_data",
	//	}
	//w.Write([]byte("cve page"))
	// TODO: в случае ошибки на нижнем уровне выводить сюда ошибку
	w.Write([]byte(CACHE.GetAllcveJson()))
	//	packet_list_Templ.Execute(w, &v)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	//	var v = struct {
	//		Host string
	//		Data string
	//	}{
	//		r.Host,
	//		"some_data",
	//	}
	w.Write([]byte("home page"))
	//	packet_list_Templ.Execute(w, &v)
}

func Start() {
	log.Println("http.ListenAndServe:8081")
	//go stop_web_server() //что бы сервер сам остановился
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/api/v1/cves", serveCves)
	//	http.HandleFunc("/ws", serveWs)
	//	Data_chan <- "\r\nsadfasdf\r\n"
	if err := http.ListenAndServe(ServAddr, nil); err != nil {
		log.Println(err)
	}
}

func Start_and_open() {
	go Start()
	//_ = exec.Command("cmd", "/c", "start", "http://localhost:8081/").Start()
	_ = exec.Command("cmd", "/c", "start", "http://localhost:8081/api/v1/cves").Start()
}
