package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
)

var logger *slog.Logger

func initLogger() (*os.File, error) {
	file, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		return nil, err
	}

	// Create JSON handler writing to the file
	handler := slog.NewJSONHandler(file, nil)
	logger = slog.New(handler)
	return file, nil
}

func main() {
	logFile, err := initLogger()
	if err != nil {
		log.Fatal("failed to init logger:", err)
	}
	defer logFile.Close()

	if err := initDB(); err != nil {
		logger.Error("db init failure", "error", err.Error())
		os.Exit(1)
	}
	defer db.Close()

	http.HandleFunc("/log", tempHandler)

	logger.Info("server starting", "port", 8080)
	if err := http.ListenAndServe("192.168.1.248:8080", nil); err != nil {
		logger.Error("server crash", "error", err.Error())
		os.Exit(1)
	}
}
