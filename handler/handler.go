package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/lateralusd/laserver/db"
)

func NewHandler(db *db.DB) *Handler {
	return &Handler{
		db: db,
	}
}

type Handler struct {
	db *db.DB
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uuid := r.URL.Query().Get("id")
	if uuid != "" {
		log.Printf("Got request from %s with id %s", r.RemoteAddr, uuid)
		d := db.Data{
			UUID: uuid,
			Time: time.Now(),
		}
		found, err := h.db.Exists(uuid)
		if err != nil {
			log.Printf("Error checking for uuid presence: %v", err)
			w.WriteHeader(http.StatusOK)
			return
		}
		if !found {
			if err := h.db.Save(&d); err != nil {
				log.Printf("Error saving to db: %v", err)
				w.WriteHeader(http.StatusOK)
				return
			}
			log.Printf("%s saved successfully", uuid)
		}
		w.WriteHeader(http.StatusOK)
		return
	}
	log.Printf("Got request from %s without id", r.RemoteAddr)
	w.WriteHeader(http.StatusOK)
}
