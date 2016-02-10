package main

import (
	"log"
	"net/http"
	"os"
)

const url  = "http://10.254.253.252:2375/"
//const url  = "http://192.168.0.254:2375/"

var host,_ = os.Hostname()
var baseUrl  = "https://" + host + "/"

func main() {

	server()
}


func server(){
	http.Handle("/app/", http.FileServer(http.Dir("./appweb")))
	//http.HandleFunc("/containers",containers)
	http.HandleFunc("/",containers)

	http.HandleFunc("/containers/inspect/", containersInspect)
	http.HandleFunc("/containers/restart/", containersRestart)
	http.HandleFunc("/containers/start/", containersStart)
	http.HandleFunc("/containers/stop/", containersStop)
	http.HandleFunc("/containers/pause/", containersPause)
	http.HandleFunc("/containers/unpause/", containersUnpause)
	http.HandleFunc("/containers/delete/", containersDelete)
	log.Println("Serving at "+host+":443...")
	log.Fatal(http.ListenAndServeTLS(":443", "certificate/server.cert", "certificate/server.key", nil))

}
