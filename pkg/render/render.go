package render

// All the files for a package must exist on the same directory
import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// RenderTemplate renders a specific html template to the writer w with filename ending in tmpl.
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// Create a template cache.
	templateCache, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	// Get template from cache.
	templ, templateFound := templateCache[tmpl]
	if !templateFound {
		log.Fatal(err)
	}

	// This buffer is arbitrary code to add more error checking below. Not needed.
	// Without this, you can only error check when writing to w.
	buf := new(bytes.Buffer)
	err = templ.Execute(buf, nil)
	if err != nil {
		fmt.Println("error executing template:", err)
	}

	// Render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{} // Empty map, same as the make function.

	// Get all the files name *.fileName.gohtml from ./templates (FULL PATH from root project)
	// Ex: templates/home.page.gohtml
	fileNames, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		return templateCache, err
	}

	// Range through the received files.
	for _, fileName := range fileNames {
		name := filepath.Base(fileName) // Filename MINUS the full path. Ex home.fileName.gohtml.

		// Create template from the current fileName being iterated by the loop.
		// The line below creates a template with the name I want (home.filename.gohtml for example).
		// And it also populates it with the corresponding file.
		templateSet, err := template.New(name).ParseFiles(fileName) // Page is the FULL PATH.
		if err != nil {
			return templateCache, err
		}

		// Search for layouts, that were not searched in the initial filepath.Glob and need to be applied in
		// each template (since they are a layout).
		fileNames, err := filepath.Glob("./templates/*layout.gohtml")
		if err != nil {
			return templateCache, err
		}
		// If any layout fileNames were found, add the layout to the existing templateSet. So it mixes layout + page.
		if len(fileNames) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return templateCache, err
			}
		}

		templateCache[name] = templateSet
	}
	return templateCache, nil
}
