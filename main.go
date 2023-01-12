package main

import (
	"fmt"
	"log"
	"simple-api/routes"
	"simple-api/store"

	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()

	store.NewDb()
	// storage.MinioConnection()
}

func main() {
	app := routes.New()
	log.Fatal(app.Listen(fmt.Sprintf(":%d", 3000)))
}
