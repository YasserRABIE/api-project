package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("env files can't ne loaded due to: ", err)
	}
}

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}
	defer store.db.Close()

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(os.Getenv("PORT"), store)
	server.run()
}
