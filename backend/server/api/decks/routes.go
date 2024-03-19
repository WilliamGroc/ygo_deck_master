package decks

import (
	"ygocarddb/server/authentication"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

type DecksRoutes struct {
	DB *mongo.Database
}

type DecksRoutesInterface interface {
	InjectRoutes() *chi.Mux
}

func NewRouter(db *mongo.Database) *DecksRoutes {
	return &DecksRoutes{DB: db}
}

func (t *DecksRoutes) InjectRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Get("/", t.ListDecks)
		r.Get("/{id}", t.GetDeck)
	})

	r.Group(func(r chi.Router) {
		r.Use(authentication.TokenVerifyMiddleWare)
		r.Post("/", t.CreateDeck)
		r.Put("/{id}", t.UpdateDeck)
		r.Delete("/{id}", t.DeleteDeck)
	})

	return r
}
