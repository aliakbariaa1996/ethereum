package server

import (
	"context"
	"github.com/aliakbariaa1996/ethereum/config"
	rg "github.com/aliakbariaa1996/ethereum/internal/api/v1"
	httpx "github.com/aliakbariaa1996/ethereum/internal/http"
	"github.com/aliakbariaa1996/ethereum/internal/services/ethereum"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	*echo.Echo
	httpServer *http.Server
	logger  *log.Logger
	ss      *ServiceStorage
	cfg     *config.Config
	handler Handler
}

type ServiceStorage struct {
	etherService ethereum.UseService
}

type Handler struct {
	logger *log.Logger
}

func NewServer(cfg *config.Config, logger *log.Logger) (*Server, error) {
	var err error
	s := &Server{
		Echo:   echo.New(),
		cfg:    cfg,
		logger: logger,
	}
	s.handler = Handler{logger: s.logger}
	return s, err
}

func (srv *Server) RunServer(port string) error {
	// HTTP Server
	router := httpx.InitRouter()
	ss:=NewServiceStorage(srv.cfg,srv.logger)
	rg.Routes(router, ss.etherService)

	srv.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := srv.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return srv.httpServer.Shutdown(ctx)
}
func NewServiceStorage(cfg *config.Config, logger *log.Logger) *ServiceStorage {
	return &ServiceStorage{
		etherService:        ethereum.NewEthereumUseCase(),
	}
}