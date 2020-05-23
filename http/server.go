package http

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nhannhan159/weather-app-go/domain"

	"github.com/gin-gonic/gin"
	"github.com/nhannhan159/weather-app-go/http/middleware"
)

type httpManager struct {
	engine       *gin.Engine
	config       domain.ServerConfig
	handlers     []domain.IRootHandler
	appResources *domain.Resources
}

func NewHTTPManager(config domain.ServerConfig) domain.IHttpManager {
	gin.DisableConsoleColor()

	f, _ := os.OpenFile("gin.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	gin.DefaultWriter = io.MultiWriter(os.Stdout, f)

	return &httpManager{
		config:   config,
		engine:   gin.Default(),
		handlers: []domain.IRootHandler{},
	}
}

func (server *httpManager) RegisterHandler(handler domain.IRootHandler) {
	server.handlers = append(server.handlers, handler)
}

func (server *httpManager) RegisterResources(resources *domain.Resources) {
	server.appResources = resources
}

func (server *httpManager) Run() {
	// Initialize templates
	// router.LoadHTMLGlob("template/*")

	// Initialize middleware
	server.engine.Use(middleware.ResourcesMiddleware(server.appResources))
	server.engine.Use(middleware.LoggerMiddleware())

	// Initialize handlers
	for _, rootHandler := range server.handlers {
		rootHandler.Handle(server.engine)
	}

	// Initialize server
	srv := &http.Server{
		Addr:    ":" + server.config.Port,
		Handler: server.engine,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
