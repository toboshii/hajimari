package main

import (
	"errors"
	"net/http"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/ullbergm/hajimari/internal/config"
	"github.com/ullbergm/hajimari/internal/handlers"
	"github.com/ullbergm/hajimari/internal/log"
)

var (
	logger = log.New()
)

func init() {
	viper.SetConfigName("config")          // name of config file (without extension)
	viper.AddConfigPath("/config")         // path to look for the config file in
	viper.AddConfigPath("$HOME/.hajimari") // call multiple times to add many search paths
	viper.AddConfigPath(".")               // optionally look for config in the working directory
	viper.AutomaticEnv()

	err := viper.ReadInConfig() // Find and read the config file

	if err != nil { // Handle errors reading the config file
		panic(errors.New("Fatal error config file: " + err.Error()))
	}

	logger.Info("Using config file: ", viper.ConfigFileUsed())

	viper.OnConfigChange(func(e fsnotify.Event) {
		logger.Info("Reloading config: ", e.Name)
	})

	viper.WatchConfig()

	config.SetDefaults()
}

func main() {

	// appConfig, err := config.GetConfig()
	// if err != nil {
	// 	logger.Fatal("Failed to read configuration for hajimari", err)
	// 	return
	// }

	httpHandler := handlers.NewHandler()

	logger.Printf("Listening on :%d", 3000)
	logger.Fatal(http.ListenAndServe(":3000", httpHandler))
}
