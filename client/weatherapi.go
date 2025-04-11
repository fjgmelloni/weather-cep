package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/fjgmelloni/weather-cep/model"
)

func FetchTemperature(city string) (float64, error) {
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		return 0, fmt.Errorf("WEATHER_API_KEY não encontrada")
	}
	query := url.QueryEscape(city)

	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, query)
	fmt.Println("Requisição para WeatherAPI:", url)

	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("erro ao fazer requisição: %w", err)
	}
	defer resp.Body.Close()
	fmt.Println("Status code:", resp.StatusCode)
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Body da resposta:", string(body))

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("erro na resposta da WeatherAPI")
	}

	var data model.WeatherAPIResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return 0, fmt.Errorf("erro ao decodificar JSON: %w", err)
	}

	return data.Current.TempC, nil
}
