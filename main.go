package main

import (
	"LibraryAPI-GitFlow/src/config"
	"fmt"
	"LibraryAPI-GitFlow/src/router"
	"log"
	"net/http"
)

func main(){
	config.Initialize()
	r := router.Create()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}