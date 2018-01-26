package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/leninhasda/gitpull-me/api"
)

func main() {
	router := api.Router()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3009"
	}

	srvr := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGKILL, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		log.Println("Server started at:", port)
		log.Fatal(srvr.ListenAndServe())
	}()

	<-stop
}
