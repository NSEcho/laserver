package handler

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/lateralusd/laserver/db"
)

func NewHandler(db *db.DB, logPath string) *Handler {
	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	wr := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wr)
	return &Handler{
		db: db,
		f:  f,
	}
}

type Handler struct {
	db *db.DB
	f  *os.File
}

func (h *Handler) Close() {
	h.f.Close()
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
