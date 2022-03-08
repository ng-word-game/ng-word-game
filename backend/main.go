package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	addr = flag.String("addr", "0.0.0.0:443", "http service address")
)

func main() {
	flag.Parse()

	mux := http.NewServeMux()
	wsHandler := NewWshandler()
	mux.HandleFunc("/", wsHandler.ServeWebsocket)
	mux.HandleFunc("/rooms", wsHandler.GetRooms)

	log.Fatal(http.ListenAndServe(*addr, mux))
}
