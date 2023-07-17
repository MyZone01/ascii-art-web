package main

import (
	ascii "ascii_art_web/lib"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const PORT = ":8080"

func validateRequest(req *http.Request, res http.ResponseWriter, url, method string) bool {
	if req.URL.Path != url {
		res.WriteHeader(http.StatusNotFound)
		renderTemplate(res, "404")
		log.Println("404 ❌ - Page not found ")
		return false
	}

	if req.Method != method {
		res.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(res, "%s", "Error - Method not allowed")
		log.Println("405 ❌ - Method not allowed")
		return false
	}
	return true
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	if validateRequest(req, res, "/", http.MethodGet) {
		renderTemplate(res, "index")
		log.Println("200 ✅")
	}
}

func asciiHandler(res http.ResponseWriter, req *http.Request) {
	if validateRequest(req, res, "/ascii-art", http.MethodPost) {
		text := strings.ReplaceAll(req.FormValue("text"), "\r\n", "\n")
		font := req.FormValue("font")
	
		if font != "standard" && font != "shadow" && font != "thinkertoy" {
			res.WriteHeader(http.StatusBadRequest)
			log.Println("400 ❌ Bad Request - Font Not Found")
			fmt.Fprintf(res, "%s", "Error - Font Not Found")
			return
		}
		asciiCharacters, err := ascii.ParseFile("fonts/"+font+".txt", false)
		if err {
			res.WriteHeader(http.StatusInternalServerError)
			log.Println("500 ❌ Internal Server Error - Template file not found")
			fmt.Fprintf(res, "%s", "Error - Font Not Found")
			return
		}
	
		output, err := ascii.ConvertTextToArt(text, "left", "", "", asciiCharacters)
		if err {
			res.WriteHeader(http.StatusBadRequest)
			log.Println("400 ❌ Bad Request - Non valid character")
			fmt.Fprintf(res, "%s", "Error - Non valid character")
			return
		}
		fmt.Fprintf(res, "%s", output)
		log.Println("200 ✅ =>", req.FormValue("text"))
	}
}

func exportFile(res http.ResponseWriter, req *http.Request) {
	if validateRequest(req, res, "/download", http.MethodGet) {
		output := req.FormValue("output")
	
		res.Header().Set("Content-Type", "text/plain")
		res.Header().Set("Content-Length", strconv.Itoa(len(output)))
		res.Header().Set("Content-Disposition", "attachement; filename=file.txt")
		res.Write([]byte(output))
	}
}

var templates = template.Must(template.ParseGlob("./template/*.html"))

func renderTemplate(res http.ResponseWriter, tmpl string) {
	err := templates.ExecuteTemplate(res, tmpl+".html", nil)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		fmt.Fprintf(res, "%s", "Error - Cannot render page")
		log.Println("500 ❌ Internal Server Error")
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ascii-art", asciiHandler)
	http.HandleFunc("/download", exportFile)

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./static/css/"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./static/img/"))))

	fmt.Println("Server started and listening on", PORT)
	fmt.Println("http://localhost" + PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
