package main

import (
	"os"

	"github.com/Nie-Mand/go-api/internal"
	"github.com/Nie-Mand/go-api/internal/utils/log"
	"go.uber.org/zap"
	// _ "github.com/joho/godotenv/autoload"
)

func main() {
	err := log.InitLogger("debug", "dev")
	handleError(err)

	internal.Run()
}

func handleError(err error) {
	if err != nil {
		zap.L().Error(err.Error())
		os.Exit(1)
	}
}
