package modules

import (
	"fmt"

	"github.com/tidwall/gjson"
	"github.com/toboshii/hajimari/internal/log"
)

var (
	logger = log.New()
)

// GetWeather fetches weather information from http://wttr.in
func GetWeather(data map[string]string) string {
	body, err := fetchWeather("http://wttr.in/" + data["city"] + "?format=j1")
	if err != nil {
		return ""
	}

	info := gjson.GetManyBytes(body, "current_condition.0.temp_F", "current_condition.0.weatherDesc.0.value")
	if info[0].String() == "" || info[1].String() == "" {
		return ""
	}

	return fmt.Sprintf("Temperature in %s: %s degrees, current condition is: %s %s",
		data["city"],
		info[0].String(),
		info[1].String(),
		updatedAt(),
	)
}

var fetchWeather = actualFetch
