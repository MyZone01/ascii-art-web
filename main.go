package main

import (
	ascii "ascii_art_web/lib"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

const PORT = ":8080"

func indexHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		res.WriteHeader(http.StatusNotFound)
		renderTemplate(res, "error")
		log.Println("404 ❌ - Page not found ")
		return
	}

	if req.Method != http.MethodGet {
		res.WriteHeader(http.StatusMethodNotAllowed)
		renderTemplate(res, "error")
		log.Println("405 ❌ - Method not allowed")
		return
	}
	renderTemplate(res, "index")
	log.Println("200 ✅")
}

func asciiHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/ascii-art" {
		res.WriteHeader(http.StatusNotFound)
		renderTemplate(res, "error")
		log.Println("404 ❌ - Page not found ")
		return
	}

	if req.Method != http.MethodPost {
		res.WriteHeader(http.StatusMethodNotAllowed)
		renderTemplate(res, "error")
		log.Println("405 ❌ - Method not allowed")
		return
	}

	text := strings.ReplaceAll(req.FormValue("text"), "\r\n", "\n")
	font := req.FormValue("font")

	if font != "standard" && font != "shadow" && font != "thinkertoy" {
		res.WriteHeader(http.StatusInternalServerError)
		log.Println("500 ❌ Internal Server Error - Font Not Found")
		renderTemplate(res, "error")
		return
	}
	asciiCharacters := ascii.ParseFile("fonts/"+font+".txt", false)
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

	fmt.Println("Server started and listening on ", PORT)
	fmt.Println("http://localhost" + PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
