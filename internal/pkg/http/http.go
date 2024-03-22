package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Server serves HTTP endpoints.
type Server interface {
	Run() chan error
	Router() *fiber.App
}

type server struct {
	server   *http.Server
	fiberApp *fiber.App
	cfg      Config
}

// Config is basic HTTP server config.
type Config struct {
	Host         string
	Port         string
	IdleTimeout  time.Duration
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	Prefork      bool
}

// New to create new web server.
func New(cfg Config) Server {
	return &server{
		fiberApp: fiber.New(fiber.Config{
			IdleTimeout:  cfg.IdleTimeout,
			WriteTimeout: cfg.WriteTimeout,
			ReadTimeout:  cfg.ReadTimeout,
			Prefork:      cfg.Prefork,
		}),
		cfg: cfg,
	}
}

// Router returns server router.
func (s *server) Router() *fiber.App {
	return s.fiberApp
}

// Run to start serving HTTP.
func (s *server) Run() chan error {
	var ch = make(chan error)
	go func(ch chan error) {
		err := s.fiberApp.Listen(fmt.Sprintf("%s:%s", s.cfg.Host, s.cfg.Port))
		if err != nil {
			ch <- err
			return
		}

		s.server = &http.Server{
			ReadTimeout:  s.cfg.ReadTimeout,
			WriteTimeout: s.cfg.WriteTimeout,
		}
	}(ch)
	return ch
}
