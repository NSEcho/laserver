package main

import (
	"log"
	"net/http"

	"github.com/lateralusd/laserver/db"
	"github.com/lateralusd/laserver/handler"
)

func main() {
	db := db.NewDB("data.db")
	defer db.Close()

	h := &handler.Handler{
		DB: db,
	}
	http.Handle("/", h)
	log.Fatal(http.ListenAndServe(":3300", nil))
}
