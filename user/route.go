package user
 
import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []User
	users, err := GetUsers()
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)

	if err != nil {
        // If there's an error, set the status to 500 Internal Server Error
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/users", GetUsersHandler).Methods("GET")
}

