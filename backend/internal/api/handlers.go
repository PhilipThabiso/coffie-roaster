package api

import (
	"encoding/json"
	"net/http"

	"github.com/PhilipThabiso/coffie-roaster/internal/db"
	"github.com/PhilipThabiso/coffie-roaster/logger"
)

type Reading struct {
	Temperature float64 `json:"temperature"`
}

func TempHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		logger.Log.Warn("invalid method", "method", r.Method, "remote", r.RemoteAddr)
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var data Reading
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		logger.Log.Error("decode failure", "error", err.Error(), "remote", r.RemoteAddr)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if err := db.SaveReading(data.Temperature); err != nil {
		logger.Log.Error("db save failure", "error", err.Error(), "temp", data.Temperature)
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	logger.Log.Info("reading saved", "temp", data.Temperature, "remote", r.RemoteAddr)
	w.WriteHeader(http.StatusCreated)
}
