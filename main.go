package main

import (
	"embed"
	"errors"
	"io/fs"
	"net/http"
	"text/template"
	"time"

	loggerMiddleware "github.com/chi-middleware/logrus-logger"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
	"github.com/toboshii/hajimari/internal/handlers"
	"github.com/toboshii/hajimari/internal/log"
	"github.com/toboshii/hajimari/internal/services"
	"github.com/toboshii/hajimari/internal/stores"
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

	// appConfig, err := config.GetConfig()
	// if err != nil {
	// 	logger.Fatal("Failed to read configuration for hajimari", err)
	// 	return
	// }

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(loggerMiddleware.Logger("router", logger))
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	staticFiles, _ := fs.Sub(staticFiles, "web")
	r.Handle("/static/*", http.FileServer(http.FS(staticFiles)))

	service := services.NewStartpageService(stores.NewFileStore(), logger)

	r.Mount("/", handlers.NewStartpageResource(service, tpl).StartpageRoutes())

	logger.Printf("Listening on :%d\n", 3000)
	logger.Fatal(http.ListenAndServe(":3000", r))
}
