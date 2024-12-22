package main

import (
	"os"

	"github.com/Nie-Mand/go-api/internal"
	"github.com/Nie-Mand/go-api/internal/utils/log"
	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap"
)

func main() {
	err := log.InitLogger("debug", "dev")
	handleError(err)
	handleError(internal.Run())
}

func handleError(err error) {
	if err != nil {
		zap.L().Error(err.Error())
		os.Exit(1)
	}
}
