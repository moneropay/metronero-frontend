package utils

import (
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

func StartServerWithGracefulShutdown(a *fiber.App, bind string) {
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := a.Shutdown(); err != nil {
			log.Println("Server is not shutting down: ", err)
		}

		close(idleConnsClosed)
	}()

	if err := a.Listen(bind); err != nil {
		log.Println("Failed to start server: ", err)
	}

	<-idleConnsClosed
}
