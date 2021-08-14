package modules

import (
	"fmt"

	"github.com/tidwall/gjson"
)

// GetPiholeStats fetches basic statistics from Pi-hole's API
func GetPiholeStats(data map[string]string) string {
	body, err := fetchPihole(data["summaryEndpoint"])
	if err != nil {
		return ""
	}

	info := gjson.GetManyBytes(body, "status", "ads_percentage_today")

	if len(info) == 0 {
		return ""
	}

	return fmt.Sprintf("Status: %s. Ads blocked today: %.1f%% %s",
		info[0].String(),
		info[1].Float(),
		updatedAt(),
	)
}

var fetchPihole = actualFetch
