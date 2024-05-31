package config

import (
	"time"

	"github.com/spf13/viper"
	"github.com/ullbergm/hajimari/internal/models"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Config struct for hajimari
type Config struct {
	NamespaceSelector     NamespaceSelector       `json:"namespaceSelector"`
	DefaultEnable         bool                    `json:"defaultEnable"`
	InstanceName          string                  `json:"instanceName"`
	Title                 string                  `json:"title"`
	Name                  string                  `json:"name"`
	LightTheme            string                  `json:"lightTheme"`
	DarkTheme             string                  `json:"darkTheme"`
	CustomThemes          []models.Theme          `json:"customThemes"`
	ShowSearch            bool                    `json:"showSearch"`
	ShowGreeting          bool                    `json:"showGreeting"`
	ShowApps              bool                    `json:"showApps"`
	ShowAppGroups         bool                    `json:"showAppGroups"`
	ShowAppUrls           bool                    `json:"showAppUrls"`
	ShowAppInfo           bool                    `json:"showAppInfo"`
	ShowAppStatus         bool                    `json:"showAppStatus"`
	DefaultAppIcon        string                  `json:"defaultAppIcon"`
	ShowBookmarks         bool                    `json:"showBookmarks"`
	ShowGlobalBookmarks   bool                    `json:"showGlobalBookmarks"`
	ShowBookmarkGroups    bool                    `json:"showBookmarkGroups"`
	AlwaysTargetBlank     bool                    `json:"alwaysTargetBlank"`
	DefaultSearchProvider string                  `json:"defaultSearchProvider"`
	SearchProviders       []models.SearchProvider `json:"searchProviders"`
	CustomApps            []models.AppGroup       `json:"customApps"`
	GlobalBookmarks       []models.BookmarkGroup  `json:"globalBookmarks"`
	Modules               []Module                `json:"modules"`
	Experimental          []ExperimentalFeature
}

type Module struct {
	Name           string
	UpdateInterval time.Duration
	Data           map[string]string
	Output         string
}

// NamespaceSelector struct for selecting namespaces based on labels and names
type NamespaceSelector struct {
	Any           bool
	MatchNames    []string
	LabelSelector *metav1.LabelSelector
}

// ExperimentalFeature struct for featureGating new experiments
type ExperimentalFeature struct {
	Enabled    bool
	Name       string
	Properties map[string]bool
}

func SetDefaults() {
	viper.SetDefault("DefaultEnable", false)
	viper.SetDefault("Title", "Hajimari")
	viper.SetDefault("Name", "You")
	viper.SetDefault("LightTheme", "gazette")
	viper.SetDefault("DarkTheme", "horizon")
	viper.SetDefault("CustomThemes", []models.Theme{})
	viper.SetDefault("ShowSearch", true)
	viper.SetDefault("ShowGreeting", true)
	viper.SetDefault("ShowApps", true)
	viper.SetDefault("ShowAppGroups", false)
	viper.SetDefault("ShowAppUrls", true)
	viper.SetDefault("ShowAppInfo", false)
	viper.SetDefault("ShowAppStatus", true)
	viper.SetDefault("DefaultAppIcon", "mdi:application")
	viper.SetDefault("ShowBookmarks", true)
	viper.SetDefault("ShowGlobalBookmarks", false)
	viper.SetDefault("ShowBookmarkGroups", true)
	viper.SetDefault("AlwaysTargetBlank", false)
	viper.SetDefault("DefaultSearchProvider", "Google")
	viper.SetDefault("SearchProviders", []models.SearchProvider{
		{
			Name:      "Google",
			Token:     "g",
			Icon:      "simple-icons:google",
			SearchUrl: "https://www.google.com/search?q={query}",
			URL:       "https://www.google.com",
		},
		{
			Name:      "DuckDuckGo",
			Token:     "d",
			Icon:      "simple-icons:duckduckgo",
			SearchUrl: "https://duckduckgo.com/?q={query}",
			URL:       "https://duckduckgo.com",
		},
		{
			Name:      "IMDB",
			Token:     "i",
			Icon:      "simple-icons:imdb",
			SearchUrl: "https://www.imdb.com/find?q={query}",
			URL:       "https://www.imdb.com",
		},
		{
			Name:      "Reddit",
			Token:     "r",
			Icon:      "simple-icons:reddit",
			SearchUrl: "https://www.reddit.com/search?q={query}",
			URL:       "https://www.reddit.com",
		},
		{
			Name:      "YouTube",
			Token:     "y",
			Icon:      "simple-icons:youtube",
			SearchUrl: "https://www.youtube.com/results?search_query={query}",
			URL:       "https://www.youtube.com",
		},
		{
			Name:      "Spotify",
			Token:     "s",
			Icon:      "simple-icons:spotify",
			SearchUrl: "https://open.spotify.com/search/{query}",
			URL:       "https://open.spotify.com",
		},
	})
	viper.SetDefault("GlobalBookmarks", []models.BookmarkGroup{})
}

// GetConfig returns hajimari configuration
func GetConfig() (*Config, error) {
	var c Config

	err := viper.Unmarshal(&c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
