package main

import (
	"log"
	"HelloGo/src/cmd/k8s"
	"net/http"
	"os"
)

func main() {

	log.Println("Starting the service ...")

	port:=os.Getenv("PORT")
	if port=="" {
		log.Fatal("Port is not set.")
	}

	r:=k8s.Router()

	log.Println("the service is ready to listen and serve.")
    log.Println(port)
	log.Fatal(http.ListenAndServe(port,r))

}