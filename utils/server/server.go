package server

import (
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"

	"gitlab.com/moneropay/metronero-frontend/utils/config"
)

type Server struct {
	*fiber.App
	config.Config
}

func Init(cfg config.Config) *Server {
	engine := html.New("./views", ".html")

	if cfg.Debug {
		engine.Reload(true)
		engine.Debug(true)
	}

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	return &Server{app, cfg}
}

func (s *Server) StartServerWithGracefulShutdown() {
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := s.Shutdown(); err != nil {
			log.Println("Server is not shutting down: ", err)
		}

		close(idleConnsClosed)
	}()

	if err := s.Listen(s.Config.Bind); err != nil {
		log.Println("Failed to start server: ", err)
	}

	<-idleConnsClosed
}
