package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"main/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Post("/upload", handlers.Upload)
	app.Get("/download", handlers.Download)

	go func() {
		if err := app.Listen("127.0.0.1:8080"); err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Println("Start graceful shutdown")

	err := app.Shutdown()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Fiber server exited properly")
	}
}
