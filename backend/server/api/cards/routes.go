package cards

import (
	"ygocarddb/server/authentication"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

type CardsRoutes struct {
	DB *mongo.Database
}

type CardsRoutesInterface interface {
	InjectRoutes() *chi.Mux
}

func NewRouter(db *mongo.Database) *CardsRoutes {
	return &CardsRoutes{DB: db}
}

func (t *CardsRoutes) InjectRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Get("/{id}", t.GetCard)
		r.Get("/{id}/image", t.GetCardImage)
		r.Get("/{id}/image/big", t.GetCardImageBig)
	})

	r.Group(func(r chi.Router) {
		r.Use(authentication.TokenVerifyMiddleWare)
		r.Post("/load", t.LoadCards)
	})

	return r
}
