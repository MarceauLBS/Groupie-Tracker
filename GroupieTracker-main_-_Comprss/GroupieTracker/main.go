package main

import (
	"GroupieTracker/controller"
	"GroupieTracker/model"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"
)

var tableau [52]*model.Artist

func makeTab() {
	start := time.Now()
	for i := 0; i < 52; i++ {
		tableau[i] = controller.GetDataByID(i, false)
	}
	elapsed := time.Since(start)
	fmt.Printf("\ntook %s for array initialization\n", elapsed)

}

func join(s ...string) string {
	return strings.Join(s[1:], s[0])
}

func replace(input, from, to string) string {
	return strings.Replace(input, from, to, -1)
}

func mainPageHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" || r.Method != "GET" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	tmpl, err := template.New("index.html").Funcs(template.FuncMap{"join": join}).ParseFiles("index.html")

	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := tmpl.Execute(w, tableau); err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func artistsPageHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Path[len("/artist/"):]

	idArtist, err := strconv.Atoi(id)
	if err != nil {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	if idArtist > 51 {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	artist := controller.GetDataByID(idArtist, false)

	tmpl, err := template.New("ArtistsDetails.html").Funcs(template.FuncMap{"join": join, "replace": replace}).ParseFiles("./assets/pages/ArtistsDetails.html")

	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := tmpl.Execute(w, artist); err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		t, err := template.ParseFiles("./assets/pages/error.html")
		if err != nil {
			log.Fatal(err.Error())
		}

		err = t.Execute(w, "404 Not Found")
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

func main() {
	controller.Init()
	makeTab()
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", mainPageHandler)
	http.HandleFunc("/artist/", artistsPageHandler)
	http.ListenAndServe(":8080", nil)
}
