package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	addr = flag.String("addr", "0.0.0.0:443", "http service address")
)

func main() {
	flag.Parse()
	loggingSettings("log")

	mux := http.NewServeMux()
	wsHandler := NewWshandler()
	mux.HandleFunc("/", wsHandler.ServeWebsocket)
	mux.HandleFunc("/rooms", wsHandler.GetRooms)

	log.Fatal(http.ListenAndServe(*addr, mux))
}

func loggingSettings(filename string) {
	if f, err := os.Stat("./.logs"); os.IsNotExist(err) || !f.IsDir() {
		os.Mkdir(".logs", 0777)
	}
	logfile, err := os.OpenFile("./.logs/" + filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
	}
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetOutput(multiLogFile)
}
