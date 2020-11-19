package server

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type internalError struct {
	Message string `json:"message"`
}

func SetupApp() {
	app := fiber.New(fiber.Config{
		Concurrency:  	256 * 1024,
		WriteTimeout: 	10 * time.Second,
		ReadTimeout: 	10 * time.Second,
		IdleTimeout:	10 * time.Second,
		BodyLimit:		4 * 1024 * 1024,
		CompressedFileSuffix: ".fiber.gz",
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError
			var msg string
			// Retrieve the custom status code if it's an fiber.*Error
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
				msg = e.Message
			}

			if msg == "" {
				msg = "cannot process the http call"
			}

			// Send custom error page
			err = ctx.Status(code).JSON(internalError{
				Message: msg,
			})
			return nil
		},
	})

	app.Use(limiter.New(limiter.Config{
		Expiration: 10 * time.Second,
		Max:      14,
	}))

	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2000",
		TimeZone:   "America/Sao_Paulo",
		Output:     os.Stdout,
	}))

	setupRoutes(app)

	// port := os.Getenv("APP_PORT")
	sslport := os.Getenv("APP_SSL_PORT")

	log.Printf( "Server up and running: https://127.0.0.1:%s", sslport)
	log.Fatal(app.Server().ListenAndServeTLS(":"+sslport, "./certs/server.crt", "./certs/server.key"))
	// log.Fatal(app.Listen(":"+port))

}

func botHealthCheck(c *fiber.Ctx) error {
	return c.SendString("Pong")
}

func setupRoutes(app *fiber.App) {
	app.Get("/ping", botHealthCheck)
}
