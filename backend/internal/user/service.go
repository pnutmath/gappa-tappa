package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Service ...
type Service struct {
	repo *Repository
}

// NewUserService ...
func NewUserService(repo *Repository) *Service {
	return &Service{repo}
}

// Get ...
func (srv *Service) Get(w http.ResponseWriter, r *http.Request) {

	user, err := srv.repo.Get(mux.Vars(r)["username"])
	if err != nil {
		fmt.Fprintf(w, "%v", err)
	}
	userdata, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%v", string(userdata))
}
