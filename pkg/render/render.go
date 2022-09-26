package render

// All the files for a package must exist on the same directory
import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templateCache = make(map[string]*template.Template)

// RenderTemplate renders a specific html template to the writer w with filename ending in tmpl.
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// ParseFiles uses the root of the application as the base string.
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.gohtml")
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

func RenderTemplateFromCache(w http.ResponseWriter, templateName string) {
	var tmpl *template.Template
	var err error

	// Check to see if we have the template in our cache and return it. Otherwise, add it to the cache.
	_, inMap := templateCache[templateName]
	if !inMap {
		log.Println("adding template to cache")
		err = createTemplateCache(templateName)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println("using cached template")
	}

	tmpl = templateCache[templateName]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(templateName string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", templateName),
		"./templates/base.layout.gohtml",
	}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	templateCache[templateName] = tmpl
	return nil
}
