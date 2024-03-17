package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// Set time to UTC+7
	jktLocation, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Fatal(err)
	}
	time.Local = jktLocation

	if len(os.Args) < 1 {
		// Load env
		err = godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}
}
