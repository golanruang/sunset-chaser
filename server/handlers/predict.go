package handlers

import (
	"net/http"
	"sunset-chaser/types"

	"github.com/gin-gonic/gin"
)

func HandleSunsetPrediction(c *gin.Context) {
	var request types.PredictionRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request format",
		})
		return
	}

	weatherData, err := getCurrentWeather(request.Latitude, request.Longitude)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve weather data",
		})
		return
	}

	// (Optional) Add a basic score prediction
	// score := scoreSunset(weatherData)

	c.JSON(http.StatusOK, gin.H{
		"weatherData": weatherData,
		// "score":   score,
	})
}

func getCurrentWeather(latitude, longitude float64) (types.WeatherData, error) {
	return types.WeatherData{
		CloudCover:  0.3,
		Humidity:    65.0,
		WindSpeed:   8.0,
		Visibility:  10.0,
		Temperature: 72.0,
		Description: "Partly cloudy",
	}, nil
}

func calculateSunsetScore(weather types.WeatherData) float64 {
	score := 0
	return score
}
