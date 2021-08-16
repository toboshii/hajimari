package main

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"text/template"

	"github.com/spf13/viper"
	"github.com/toboshii/hajimari/internal/config"
	"github.com/toboshii/hajimari/internal/handlers"
	"github.com/toboshii/hajimari/internal/log"
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
	if err != nil {             // Handle errors reading the config file
		panic(errors.New("Fatal error config file: " + err.Error()))
	}
}

//go:embed web/template
var indexHTML embed.FS

//go:embed web/static
var staticFiles embed.FS

var tpl = template.Must(template.ParseFS(indexHTML, "web/template/index.html.tmpl"))

func main() {

	appConfig, err := config.GetConfig()
	if err != nil {
		logger.Fatal("Failed to read configuration for hajimari", err)
		return
	}

	staticFiles, _ := fs.Sub(staticFiles, "web")
	http.Handle("/static/", http.FileServer(http.FS(staticFiles)))

	page := handlers.NewPage(appConfig, tpl)
	http.HandleFunc("/", page.FrontpageHandler)

	logger.Printf("Listening on :%d\n", 3000)
	logger.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 3000), nil))
}
