package main

import (
	"log"
	"net/http"
	"os"

	"github.com/PhilipThabiso/coffie-roaster/internal/api"
	"github.com/PhilipThabiso/coffie-roaster/internal/db"
	"github.com/PhilipThabiso/coffie-roaster/logger"
)

func main() {
	logFile, err := logger.Init()
	if err != nil {
		log.Fatal("failed to init logger.log:", err)
	}
	defer logFile.Close()

	if err := db.InitDB(); err != nil {
		logger.Log.Error("db init failure", "error", err.Error())
		os.Exit(1)
	}
	defer db.DB.Close()

	http.HandleFunc("/log", api.TempHandler)

	logger.Log.Info("server starting", "port", 8080)
	if err := http.ListenAndServe("192.168.1.248:8080", nil); err != nil {
		logger.Log.Error("server crash", "error", err.Error())
		os.Exit(1)
	}
}
