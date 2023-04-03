package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"timestamps-assignment/api/app"
	h "timestamps-assignment/api/http"
)

func main() {
	timestampsService := app.NewTimestampDataService()
	timestampsHandler := h.NewTimestampDataHandler(*timestampsService)
	timestampsRouter := h.NewTimestampsRoutes(*timestampsHandler)
	mainRouter := h.InitMainRouter(&timestampsRouter)
	muxRouter := mainRouter.InitRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), muxRouter))
}
