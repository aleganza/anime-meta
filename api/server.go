package main

import (
	"encoding/json"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/api/users", usersHandler)
	http.ListenAndServe(":" + os.Getenv("PORT"), nil)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	users := []string{"Ale", "Mario"}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(users)
}