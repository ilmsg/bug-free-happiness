package main

import (
	"log/slog"
	"os"
)

func main() {
	slog.Info("hello, world!", "user", os.Getenv("USER"))

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("hello, world", "user", os.Getenv("USER"))
}
