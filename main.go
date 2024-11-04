package main

import (
	"gate_bill/db"
	"gate_bill/user"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	db.Connect()

	defer db.Conn.Close()
	router := mux.NewRouter()

	user.RegisterRoutes(router)

	fmt.Println("Server is running on port 6969...")

	err := http.ListenAndServe(":6969", router)
	if err != nil {
		panic("can't listen on port 6969")
	}
}

