package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/alexnel24/lead-collection-personal/src/app"
)

func main() {

	to_display := os.Getenv("PORT")
	fmt.Println("\nWorking on Port", to_display)
	
	server := app.NewServer()
	
	http.ListenAndServe(":"+os.Getenv("PORT"), server)
}