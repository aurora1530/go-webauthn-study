package main

import (
	"net/http"
	"os"

	"github.com/aurora1530/go-webauthn-study/backend/internal"
	"github.com/joho/godotenv"
)

func init(){
	prod := os.Getenv("PRODUCTION")
	if prod != "true" {
		err := godotenv.Load("../.env")
		if err != nil {
			panic(err)
		}
	}
}

func main(){
	server := internal.NewServer()
	router := internal.CreateRouter(server)

	http.ListenAndServe(":5000", router)
}