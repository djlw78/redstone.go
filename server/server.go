package main

import (
	"fmt"
	"github.com/feeddageek/redstone.go/minecraft"
	"log"
	"net/http"
)

func ressourceHandler(w http.ResponseWriter, req *http.Request) {
	//TODO add a handler that serve static ressources
}

func loginHandler(w http.ResponseWriter, req *http.Request) {
	//TODO add a mechanism to handle authentication
}

func startHandler(w http.ResponseWriter, req *http.Request) {
	
	if i,err := minecraft.Start(jar,world); err != nil {
		fmt.Fprintf(w, "Not started : %s", err.Error())
	} else {
		go disp()
		w.Write([]byte("Started.\n"))
	}
}

func stopHandler(w http.ResponseWriter, req *http.Request) {
	if err := i.Stop(); err != nil {
		fmt.Fprintf(w, "Not stopped : %s", err.Error())
	} else {
		w.Write([]byte("Stopped.\n"))
	}
}

func disp() {
	for i.Scan() {
		log.Println(i.Text())
	}
}

var i *minecraft.Instance
var jar minecraft.Jar
var world minecraft.World

func main() {
	jar = minecraft.NewJar("minecraft_server.jar","-Xmx900M", "-Xms900M")
	world = minecraft.NewWorld("./minecraft/")
	http.HandleFunc("/res/", ressourceHandler)
	http.HandleFunc("/login/", loginHandler)
	http.HandleFunc("/start/", startHandler)
	http.HandleFunc("/stop/", stopHandler)
	log.SetFlags(0)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
