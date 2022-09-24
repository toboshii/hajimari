package config

import (
	"time"

	"github.com/creasty/defaults"
	"github.com/spf13/viper"
	"github.com/toboshii/hajimari/internal/models"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Config struct for hajimari
type Config struct {
	NamespaceSelector   NamespaceSelector      `json:"namespaceSelector"`
	DefaultEnable       bool                   `default:"false"      json:"defaultEnable"`
	InstanceName        string                 `default:""           json:"instanceName"`
	Title               string                 `default:"Hajimari"   json:"title"`
	Name                string                 `default:"You"        json:"name"`
	LightTheme          string                 `default:"gazette"    json:"lightTheme"`
	DarkTheme           string                 `default:"blackboard" json:"darkTheme"`
	CustomThemes        []models.Theme         `default:"[]"         json:"customThemes"`
	ShowGreeting        bool                   `default:"true"       json:"showGreeting"`
	ShowAppGroups       bool                   `default:"false"      json:"showAppGroups"`
	ShowAppUrls         bool                   `default:"true"       json:"showAppUrls"`
	ShowAppInfo         bool                   `default:"true"       json:"showAppInfo"`
	ShowBookmarkGroups  bool                   `default:"true"       json:"showBookmarkGroups"`
	ShowGlobalBookmarks bool                   `default:"false"      json:"showGlobalBookmarks"`
	CustomApps          []models.AppGroup      `default:"[]"         json:"customApps"`
	GlobalBookmarks     []models.BookmarkGroup `default:"[]"         json:"globalBookmarks"`
	SearchProviders     []SearchProvider       `default:"[]"         json:"searchProviders"`
	Modules             []Module               `default:"[]"         json:"modules"`
	Experimental        []ExperimentalFeature
}

type SearchProvider struct {
	Name   string
	URL    string
	Prefix string
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

// GetConfig returns hajimari configuration
func GetConfig() (*Config, error) {
	var c Config
	defaults.Set(&c)
	err := viper.Unmarshal(&c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
