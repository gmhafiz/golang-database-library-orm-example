package main

import (
	"context"
	"fmt"
	"godb/sqlc"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	driver "github.com/jmoiron/sqlx"

	"godb/config"
	"godb/middleware"
	"godb/sqlx"
)

type App struct {
	sqlx *driver.DB

	config     *config.Configuration
	router     *chi.Mux
	httpServer *http.Server
}

func main() {
	app := &App{}

	app.config = config.New()
	app.SetupDB()
	app.SetupRouter()
	app.SetupServer()
	app.Run()
}

func (a *App) Run() {
	go func() {
		log.Printf("Serving at %s", a.httpServer.Addr)
		err := a.httpServer.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 60*time.Second)
	defer shutdown()

	_ = a.sqlx.Close()
	_ = a.httpServer.Shutdown(ctx)
}

func (a *App) SetupRouter() {
	a.router = chi.NewRouter()
	a.router.Use(middleware.Json)
	a.router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"message": "endpoint not found"}`))
	})

	sqlx.Handle(a.router, a.sqlx)
	sqlc.Handle(a.router, a.sqlx)

	printAllRegisteredRoutes(a.router)
}

func (a *App) SetupDB() {
	a.sqlx = sqlx.New(a.config.DB)
}

func (a *App) SetupServer() {
	a.httpServer = &http.Server{
		Addr:           "0.0.0.0:3080",
		Handler:        a.router,
		ReadTimeout:    5*time.Second,
		WriteTimeout:   10*time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func printAllRegisteredRoutes(r *chi.Mux) {
	walkFunc := func(method string, path string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("%-7s %s\n", method, path)

		return nil
	}
	if err := chi.Walk(r, walkFunc); err != nil {
		fmt.Print(err)
	}
}
