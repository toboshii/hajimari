package config

import (
	"time"

	"github.com/spf13/viper"
	"github.com/toboshii/hajimari/internal/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	logger = log.New()
)

// Config struct for hajimari
type Config struct {
	NamespaceSelector NamespaceSelector
	DefaultEnable     bool
	Title             string
	InstanceName      string
	Name              string
	CustomApps        []CustomApp
	Groups            []Group
	Providers         []Provider
	Modules           []Module
}

// CustomApp struct for specifying apps that are not generated using ingresses
type CustomApp struct {
	Name  string
	Icon  string
	URL   string
	Group string
}

type Group struct {
	Name  string
	Links []Link
}

type Provider struct {
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

type Link struct {
	Name string
	URL  string
}

// NamespaceSelector struct for selecting namespaces based on labels and names
type NamespaceSelector struct {
	Any           bool
	MatchNames    []string
	LabelSelector *metav1.LabelSelector
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
