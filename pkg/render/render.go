package render

// All the files for a package must exist on the same directory
import (
	"fmt"
	"html/template"
	"net/http"
)

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
