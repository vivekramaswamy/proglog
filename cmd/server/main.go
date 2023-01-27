package main

import (
	"fmt"
	"github.com/vivekramaswamy/proglog/internal/server"
	"log"
)

func main() {

	srv := server.NewHttpServer(":8080")
	fmt.Println("Starting up the server")
	defer fmt.Println("Shutting down the server")
	log.Fatal(srv.ListenAndServe())

}
