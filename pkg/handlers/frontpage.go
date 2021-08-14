package handlers

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/toboshii/hajimari/pkg/config"
	"github.com/toboshii/hajimari/pkg/hajimari"
	"github.com/toboshii/hajimari/pkg/hajimari/ingressapps"
	"github.com/toboshii/hajimari/pkg/kube"
	"github.com/toboshii/hajimari/pkg/kube/util"
	"github.com/toboshii/hajimari/pkg/log"
)

var (
	logger = log.New()
)

type Page struct {
	conf *config.Config
	tpl  *template.Template
}

func NewPage(conf *config.Config, tpl *template.Template) *Page {
	// for i, module := range conf.Modules {
	// 	switch module.Name {
	// 	case "Weather":
	// 		dispatch(&conf.Modules[i], modules.GetWeather)
	// 	case "Pi-hole":
	// 		dispatch(&conf.Modules[i], modules.GetPiholeStats)
	// 	}
	// }

	return &Page{conf, tpl}
}

// func dispatch(module *config.Module, f func(map[string]string) string) {
// 	ticker := time.NewTicker(1 * time.Second)
// 	go func() {
// 		for {
// 			<-ticker.C
// 			ch := make(chan string)
// 			go func() {
// 				ticker = time.NewTicker(module.UpdateInterval * time.Minute)
// 				weather := f(module.Data)
// 				ch <- weather
// 			}()

// 			module.Output = <-ch
// 		}
// 	}()
// }

func (p *Page) FrontpageHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	kubeClient := kube.GetClient()

	var hajimariApps []hajimari.App

	appConfig, err := config.GetConfig()
	if err != nil {
		logger.Error("Failed to read configuration for forecastle", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ingressAppsList := ingressapps.NewList(kubeClient, *appConfig)

	namespaces, err := util.PopulateNamespaceList(kubeClient, appConfig.NamespaceSelector)

	if err != nil {
		logger.Error("An error occurred while populating namespaces", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("Namespaces to look for hajimari apps: ", namespaces)
	hajimariApps, err = ingressAppsList.Populate(namespaces...).Get()

	if err != nil {
		logger.Error("An error occurred while looking for forceastle apps", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "text/html")

	tplerr := p.tpl.Execute(w, struct {
		Title     string
		Greeting  string
		Date      string
		Apps      []hajimari.App
		Groups    []config.Group
		Providers []config.Provider
		Modules   []config.Module
	}{
		Title:     appConfig.Title,
		Greeting:  greet(appConfig.Name, time.Now().Hour()),
		Date:      time.Now().Format("Mon, Jan 02"),
		Apps:      hajimariApps,
		Groups:    appConfig.Groups,
		Providers: appConfig.Providers,
		Modules:   appConfig.Modules,
	})

	if tplerr != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	elapsed := time.Since(start)
	logger.Printf("Served %s in %s", r.URL.Path, elapsed)
}

// greet returns the greeting to be used in the h1 heading
func greet(name string, currentHour int) (greet string) {
	switch currentHour / 6 {
	case 0:
		greet = "Good night"
	case 1:
		greet = "Good morning"
	case 2:
		greet = "Good afternoon"
	default:
		greet = "Good evening"
	}

	if name != "" {
		return fmt.Sprintf("%s, %s!", greet, name)
	}

	return fmt.Sprintf("%s!", greet)
}
