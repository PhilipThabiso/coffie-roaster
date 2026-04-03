package main

import (
	"encoding/json"
	"net/http"
)

type Reading struct {
	Temperature float64 `json:"temperature"`
}

func tempHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		logger.Warn("invalid method", "method", r.Method, "remote", r.RemoteAddr)
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var data Reading
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		logger.Error("decode failure", "error", err.Error(), "remote", r.RemoteAddr)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if err := saveReading(data.Temperature); err != nil {
		logger.Error("db save failure", "error", err.Error(), "temp", data.Temperature)
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	logger.Info("reading saved", "temp", data.Temperature, "remote", r.RemoteAddr)
	w.WriteHeader(http.StatusCreated)
}
