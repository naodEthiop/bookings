package render

import (
	"bytes"
	"fmt"
	"github.com/naodEthiop/bookings/pkg/config"
	"github.com/naodEthiop/bookings/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

// new template sets the config for the template package.

func NewTemplate(a *config.AppConfig) {
	app = a
}
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var cache map[string]*template.Template
	if app.UseCache {
		// get the template cache from the app config
		cache = app.TemplateCache
	} else {
		cache, _ = CreateTemplateCache()
	}

	// get requested templates from cache
	templates, ok := cache[tmpl]
	if !ok {
		fmt.Println(templates)
		log.Fatal("Template not found in cache")
	}
	buf := new(bytes.Buffer) // for error checking means this returned memory location has an error best practice.
	td = AddDefaultData(td)
	_ = templates.Execute(buf, td)

	// render the templates
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing the template to browser", err)
	}
}
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	// get all the file named *page.tmpl from ./templates
	pages, err := filepath.Glob("templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		newTemplate, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			newTemplate, err = newTemplate.ParseGlob("templates/*.layout.tmpl")
			//  add them to templates
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = newTemplate
	}
	return myCache, nil
}
