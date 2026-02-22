package main

import (
	"log"

	"github.com/green-house-grow-agro/gateway/internal/config"
	"github.com/green-house-grow-agro/gateway/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	router := server.New(cfg)

	router.Run(":" + cfg.Server.Port)
}
