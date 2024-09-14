package handler

import (
	"fmt"
	"net/http"
)

// Handler es la función exportada que Vercel invoca
func Handler(w http.ResponseWriter, r *http.Request) {
	// Obtener el parámetro de query "theme"
	theme := r.URL.Query().Get("theme")
	if theme == "" {
		theme = "dark" // El valor por defecto es "dark"
	}

	// Establecer el tipo de contenido a CSS
	w.Header().Set("Content-Type", "text/css")

	// Generar el CSS dinámicamente basado en el tema
	var css string
	if theme == "dark" {
		css = `
			body {
				background-color: #333;
				color: white;
			}
			h1 {
				color: lightblue;
			}
		`
	} else {
		css = `
			body {
				background-color: white;
				color: black;
			}
			h1 {
				color: darkblue;
			}
		`
	}

	// Enviar el CSS generado al cliente
	fmt.Fprintf(w, css)
}
