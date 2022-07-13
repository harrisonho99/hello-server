package main

import (
	"log"
	"net/http"
)

type Server struct{}

func (s *Server) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte(`<h1 style="text-align:center;">hello world</h1>`))
}
func main() {
	s := new(Server)

	log.Printf("Start server")
	http.ListenAndServe(":80", s)
}
