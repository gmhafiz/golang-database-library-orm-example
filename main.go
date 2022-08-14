package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	driver "github.com/jmoiron/sqlx"
	"gorm.io/gorm"

	"godb/config"
	"godb/db/ent"
	"godb/db/ent/ent/gen"
	gormDB "godb/db/gorm"
	"godb/db/sqlboiler"
	"godb/db/sqlc"
	"godb/db/sqlx"
	"godb/db/squirrel"
	"godb/middleware"
)

type App struct {
	sqlx *driver.DB
	gorm *gorm.DB
	ent  *gen.Client

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
	_ = a.ent.Close()

	// You cannot close database connection created by gorm

	_ = a.httpServer.Shutdown(ctx)
}

func (a *App) SetupRouter() {
	a.router = chi.NewRouter()
	a.router.Use(middleware.Json)
	a.router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"message": "endpoint not found"}`))
	})

	sqlx.Register(a.router, a.sqlx, a.config.DB.Type)
	sqlc.Register(a.router, a.sqlx, a.config.DB.Type)
	squirrel.Register(a.router, a.sqlx)
	gormDB.Register(a.router, a.gorm)
	sqlboiler.Register(a.router, a.sqlx)
	ent.Register(a.router, a.ent)

	printAllRegisteredRoutes(a.router)
}

func (a *App) SetupDB() {
	a.sqlx = sqlx.New(a.config.DB)
	a.gorm = gormDB.New(a.config.DB)
	a.ent = ent.New(a.config.DB)
}

func (a *App) SetupServer() {
	a.httpServer = &http.Server{
		Addr:           "0.0.0.0:3080",
		Handler:        a.router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   60 * time.Second,
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
