package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"pr-test-4/books"
	"pr-test-4/books/delivery"
	"pr-test-4/books/repository"
	"pr-test-4/books/usecase"
	"time"
)

type App struct {
	httpServer *http.Server
	booksUC    books.UseCase
}

func NewApp() *App {
	booksRepo := repository.NewRepository()
	return &App{
		booksUC: usecase.NewBooksUseCase(booksRepo),
	}
}

func (a *App) Run(port string) error {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
		delivery.RpsLimiter,
	)

	delivery.RegisterHTTPEndpoints(router, a.booksUC)

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}
