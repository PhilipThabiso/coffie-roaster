package logger

import (
	"log/slog"
	"os"
)

var (
	Log     *slog.Logger
	logFile = "server.log"
)

func Init() (*os.File, error) {
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		return nil, err
	}
	Log = slog.New(slog.NewJSONHandler(file, nil))
	return file, nil
}
