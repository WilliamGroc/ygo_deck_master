package server

import (
	"log"
	"ygocarddb/database"
	"ygocarddb/server/middlewares"

	"ygocarddb/server/api/cards"
	"ygocarddb/server/api/decks"
	"ygocarddb/server/api/users"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	DB     *mongo.Database
	Router *chi.Mux
}

type Application interface {
	Init() error
	Run()
}

func (a *App) Run() {
	a.Init()

	a.Router.Route("/api", func(r chi.Router) {
		r.Mount("/cards", cards.NewRouter(a.DB).InjectRoutes())
		r.Mount("/decks", decks.NewRouter(a.DB).InjectRoutes())
		r.Mount("/users", users.NewRouter(a.DB).InjectRoutes())
	})
}

func (a *App) Init() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	a.DB = database.InitMongoDb()
	a.Router = chi.NewRouter()

	a.Router.Use(middlewares.CorsMiddleware())
	a.Router.Use(middleware.AllowContentType("application/json", "text/xml"))
	a.Router.Use(middleware.Logger)
	a.Router.Use(middleware.Recoverer)

	return nil
}
