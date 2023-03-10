package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/da-wit/bookings/pkg/config"
	"github.com/da-wit/bookings/pkg/models"
)

const template_dir string = "./templates/"
const page_pattern string = template_dir + "*.page.tmpl"
const layout_pattern string = template_dir + "*.layout.tmpl"

var app *config.AppConfig

// UpdateConfig sets the config for the render package
func UpdateConfig(a *config.AppConfig) {
	app = a
}

// CreateTemplateCache creates template cache
// and parse layout files if exists
func CreateTemplateCache() (map[string]*template.Template, error) {
	tc := map[string]*template.Template{}

	pages, err := filepath.Glob(page_pattern)
	if err != nil {
		return tc, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		parsedTemplate, err := template.New(name).ParseFiles(page)
		if err != nil {
			return tc, err
		}

		layouts, err := filepath.Glob(layout_pattern)
		if err != nil {
			return tc, err
		}

		if len(layouts) > 0 {
			_, err := parsedTemplate.ParseGlob(layout_pattern)
			if err != nil {
				return tc, err
			}
		}

		tc[name] = parsedTemplate
	}
	return tc, nil
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders the template into given w.
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	var err error

	// 1. Check app.UseCache to determine to use cache or not
	if !app.UseCache {
		tc, err = CreateTemplateCache()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		tc = app.TemplateCache
	}

	// 2. Get requested template from cache
	t, inMap := tc[tmpl]
	if !inMap {
		log.Fatal("Cannot get template from cache")
	}

	// 3. Render template
	buf := new(bytes.Buffer)
	if err := t.Execute(buf, *td); err != nil {
		log.Fatal(err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatal(err)
	}
}
