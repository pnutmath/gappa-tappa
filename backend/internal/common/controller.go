package common

import (
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gorilla/mux"
	"github.com/pnutmath/gappa-tappa/backend/internal/user"
	"github.com/pnutmath/gappa-tappa/backend/pkg/websocket"
)

// SetupRoute ...
func SetupRoute(router *mux.Router, db *mongo.Database) {

	pool := websocket.NewPool()

	go pool.Start()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})

	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(pool, w, r)
	})

	// API routes
	apiRouter := router.PathPrefix("/api").Subrouter()

	// User Routes
	repo := user.NewUserRepository(db)
	userService := user.NewUserService(repo)
	userAPIRouter := apiRouter.PathPrefix("/user").Subrouter()
	userAPIRouter.HandleFunc("/{username}", userService.Get).Methods("GET")
}
