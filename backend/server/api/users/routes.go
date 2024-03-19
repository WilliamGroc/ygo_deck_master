package users

import (
	"ygocarddb/server/authentication"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersRoutes struct {
	DB *mongo.Database
}

type UsersRoutesInterface interface {
	InjectRoutes() *chi.Mux
}

func NewRouter(db *mongo.Database) *UsersRoutes {
	return &UsersRoutes{DB: db}
}

func (t *UsersRoutes) InjectRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Group(func(subrouter chi.Router) {
		subrouter.Post("/login", t.Login)
		subrouter.Post("/register", t.Register)
	})

	router.Group(func(subrouter chi.Router) {
		subrouter.Use(authentication.TokenVerifyMiddleWare)

		subrouter.Get("/{id}", t.GetUser)
		subrouter.Put("/{id}", t.UpdateUser)
		subrouter.Delete("/{id}", t.DeleteUser)
	})

	return router
}
