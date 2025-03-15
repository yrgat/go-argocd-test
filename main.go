package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Merhaba Kubernetes! 🚀")
}

func main() {
	app := fiber.New()

	app.Get("/process", func(c *fiber.Ctx) error {

		return c.JSON(fiber.Map{"message": "Processed successfully"})
	})
	app.Get("/health", func(c *fiber.Ctx) error {

		return c.JSON("OK")
	})

	fmt.Println("Application listen port 4000")
	log.Fatal(app.Listen(":4000"))
}
