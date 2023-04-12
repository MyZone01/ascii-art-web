package main

import (
	ascii "ascii_art_web/lib" // importation du paquet src
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const port = ":8080"

func indexHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		renderTemplate(res, "error")
		log.Println("405 ❌ - Method not allowed")
		return
	}
	renderTemplate(res, "index")
	log.Println("200 ✅")
}

func asciiHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		renderTemplate(res, "error")
		log.Println("405 ❌ - Method not allowed")
		return
	}
	
	text := req.FormValue("text")
	font := req.FormValue("font")
	
	if font != "standard" && font != "shadow" && font != "thinkertoy" {
		log.Println("500 ❌ Internal Server Error - Font Not Found")
		renderTemplate(res, "error")
		return
	}
	asciiCharacters := ascii.ParseFile("fonts/" + font + ".txt", false)
	output := ascii.ConvertTextToArt(text, "left", "", "", asciiCharacters)
	
	fmt.Fprintf(res, "%s", output)
	log.Println("200 ✅ =>", req.FormValue("text"))
}

var templates = template.Must(template.ParseGlob("./template/*.html"))

func renderTemplate(res http.ResponseWriter, tmpl string) {
	err := templates.ExecuteTemplate(res, tmpl+".html", nil)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		renderTemplate(res, "error")
		log.Println("505 ❌ Internal Server Error - Font Not Found")
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ascii-art", asciiHandler)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "./static/favicon.ico") })

	fmt.Println("Server started and listening on ", port)
	fmt.Println("http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
