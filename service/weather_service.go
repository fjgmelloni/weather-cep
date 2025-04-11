package service

import (
	"fmt"

	"github.com/fjgmelloni/weather-cep/client"
)

type WeatherResponse struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func GetWeatherByCEP(cep string) (*WeatherResponse, error) {
	fmt.Println("Buscando cidade para o CEP:", cep)

	location, err := client.FetchCityByCEP(cep)
	if err != nil {
		fmt.Println("Erro ao buscar cidade:", err)
		return nil, err
	}
	fmt.Println("Cidade encontrada:", location)

	tempC, err := client.FetchTemperature(location)
	if err != nil {
		fmt.Println("Erro ao buscar temperatura:", err)
		return nil, err
	}
	fmt.Println("Temperatura em Celsius:", tempC)

	return &WeatherResponse{
		TempC: tempC,
		TempF: tempC*1.8 + 32,
		TempK: tempC + 273,
	}, nil
}
