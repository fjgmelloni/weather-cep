package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/fjgmelloni/weather-cep/handler"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis do sistema")
	}
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		log.Fatal("WEATHER_API_KEY não encontrada! Verifique o arquivo .env ou variável de ambiente.")
	}

	r := gin.Default()

	r.GET("/weather", handler.WeatherHandler)

	log.Println("Servidor rodando em http://localhost:8080")
	r.Run(":8080")
}
