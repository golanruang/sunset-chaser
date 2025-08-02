package types

import "time"

type WeatherData struct {
	CloudCover     float64 `json:"cloud_cover"`      // Total cloud cover % (0–100)
	HighCloudCover float64 `json:"high_cloud_cover"` // High-level clouds (e.g., cirrus)
	MidCloudCover  float64 `json:"mid_cloud_cover"`  // Mid-level clouds (e.g., altostratus)
	LowCloudCover  float64 `json:"low_cloud_cover"`  // Low-level clouds (e.g., stratus)
	Humidity       float64 `json:"humidity"`         // % (0–100)
	Precipitation  float64 `json:"precipitation"`    // mm/hr
	WindSpeed      float64 `json:"wind_speed"`       // m/s or km/h
	Visibility     float64 `json:"visibility"`       // km
	Temperature    float64 `json:"temperature"`      // °C
	AirQuality     float64 `json:"air_quality"`      // AQI (0–500), optional from separate API
}

type PredictionRequest struct {
	UserID     int       `json:"user_id"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	SunsetTime time.Time `json:"sunset_time"` // Sunset time in UTC
}
