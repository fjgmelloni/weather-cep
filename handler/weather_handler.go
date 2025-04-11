package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/fjgmelloni/weather-cep/service"
)

func WeatherHandler(c *gin.Context) {
	cep := c.Query("cep")
	if len(cep) != 8 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid zipcode"})
		return
	}

	result, err := service.GetWeatherByCEP(cep)
	if err != nil {
		if err.Error() == "can not find zipcode" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, result)
}
