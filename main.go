package main

import (
	"PhoneStreamingTest/adb"
	"PhoneStreamingTest/network"
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func main() {
	flag.Parse()
	Manager := adb.Manager{}
	Manager.Start()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		network.ServeSocket(writer, request)
	})

	log.Printf("Listening on port 8080...")

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("Listen and Serve: ", err)
	}
}
