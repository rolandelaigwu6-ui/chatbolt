package main

import (
	"http/query"
	"net/http"
	"log"
	"fmt"
)


func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", query.PingHandler)
	mux.HandleFunc("/hello", query.QueryHandler)
	mux.HandleFunc("/count", query.CountHandler)
	mux.HandleFunc("/calculate",query.CalculateHandler)
	mux.HandleFunc("/agent", query.AgentHandler)
	mux.HandleFunc("/echo", query.EchoHandler)
	port := ":8080"
	server := http.Server{
		Addr: port,
		Handler:mux ,
	}
	fmt.Println("server running on port 8080")
	log.Fatal(server.ListenAndServe())
}