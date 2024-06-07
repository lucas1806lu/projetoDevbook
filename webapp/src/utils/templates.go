package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

//CarregarTemplatesbinsere os tamplates html na valiavel templates
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
	templates = template.Must(templates.ParseGlob("views/templates/*.html"))
}

//ExecutarTemplate renderia uma pagina html na tela
func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}
