package main

import (
	"context"
	"ecommerce-service/internal/app/server/handler"
	"ecommerce-service/internal/app/server/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func gracefulShutdown(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("Gin server shutting down....")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Gin server forced to shutdown: %v\n", err)
	} else {
		log.Print("Gin server existing")
	}
}

func main() {

	// database.InitDB()
	// database.Migrate()

	h := handler.NewHandler(int32(200), int32(0))

	serviceRouter := routers.NewRouter(h.Routes())

	srv := &http.Server{
		Addr:    ":5000",
		Handler: serviceRouter,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("unable to run gin service server: %v", err)
		}
	}()

	gracefulShutdown(srv)
}
