package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/fjgmelloni/weather-cep/model"
)

func FetchCityByCEP(cep string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data model.ViaCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	if data.Localidade == "" {
		return "", errors.New("can not find zipcode")
	}

	return data.Localidade, nil
}
