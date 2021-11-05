package main

import (
	"encoding/json"
	"net/http"
)

// Atributos con primera letra en Mayuscula es público y minuscula es privado
type Curso struct {
	Title     string `json:"titulo"`
	NumVideos int    `json:"num_videos"` //Modifica como se muestra en el JSON
}

type Cursos []Curso

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		cursos := Cursos{
			Curso{"Curso de Go", 30},
			Curso{"Curso de Ruby", 20},
			Curso{"Curso de NodeJS", 40},
			Curso{"Curso de Angular", 100},
		}

		json.NewEncoder(rw).Encode(cursos)
	})

	http.ListenAndServe(":8000", nil)
}
