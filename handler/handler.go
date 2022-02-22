package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/lateralusd/laserver/db"
)

type Handler struct {
	DB *db.DB
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Print("Got request from", r.RemoteAddr)
	uuid := r.URL.Query().Get("id")
	if uuid != "" {
		d := db.Data{
			UUID: uuid,
			Time: time.Now(),
		}
		if err := h.DB.Save(&d); err != nil {
			log.Printf("Error saving to db: %v", err)
			return
		}
		log.Printf("%s saved successfully", uuid)
	}
}
