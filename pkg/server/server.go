package server

import (
	"log"
	"net/http"
	"os"

	"github.com/toaru/clean-arch-api/pkg/server/interface/api"
	"github.com/volatiletech/sqlboiler/boil"
)

func init() {
	boil.DebugMode = false
}

func RunServer() {
	// DI
	config := getConfig()
	apiHandler := api.NewHandler(&config)

	// start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on port %s", port)
	err := http.ListenAndServe(":"+port, apiHandler)
	if err != nil {
		log.Fatalln(err)
	}
}
