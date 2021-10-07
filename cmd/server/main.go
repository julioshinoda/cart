package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/julioshinoda/cart/internal/api/handler"
	"github.com/julioshinoda/cart/internal/repository"
	"github.com/julioshinoda/cart/internal/usecase/cart"
	"github.com/julioshinoda/cart/pkg/logger"
	rds "github.com/julioshinoda/cart/pkg/redis"
	"go.uber.org/fx"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	kill := make(chan os.Signal, 1)
	signal.Notify(kill)

	go func() {
		<-kill
		cancel()
	}()

	app := fx.New(
		fx.Provide(logger.NewZapLogger),
		fx.Provide(rds.NewRedisClient),
		fx.Provide(repository.NewCartRedis),
		fx.Provide(cart.NewService),
		fx.Provide(handler.NewCartHandlers),
		fx.Invoke(runHttpServer),
	)
	if err := app.Start(ctx); err != nil {
		fmt.Println(err)
	}
}

func runHttpServer(lifecycle fx.Lifecycle, cartHandler *handler.CartHandler) {
	lifecycle.Append(fx.Hook{OnStart: func(context.Context) error {
		//TODO: use ENV VAR
		addr := fmt.Sprintf(":%s", "8084")
		fmt.Printf("Starting server on %v\n", addr)
		return http.ListenAndServe(addr, router(cartHandler))
	}})
}

func router(cartHandler *handler.CartHandler) http.Handler {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.Route("/cart", func(r chi.Router) {
		handler.MakeCartHandlers(r, cartHandler)
	})

	return r
}
