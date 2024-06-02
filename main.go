package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type HelloHandler struct{}

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "HelloHandler method - using http.Handle()")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello function - using http.HandleFunc()")
}

func home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/home.html")
	if err != nil {
		panic("No such template found")
	}
	t.Execute(w, "Chary")
}

func register(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/form.html")
	if err != nil {
		panic("register err")
	}
	t.Execute(w, nil)
}

func success(w http.ResponseWriter, r *http.Request) {

	firstName := r.FormValue("firstname") // automitically calls r.ParseForm()
	lastName := r.FormValue("lastname")
	fmt.Println(r.Form) // map[firstname:[Chary] lastname:[Garry]]
	fmt.Println(r.Form["firstname"])

	info := map[string]string{
		"last_name":  lastName,
		"first_name": firstName,
	}

	bson, err := json.Marshal(info)
	if err != nil {
		panic("Could not parse to json")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(bson)
}

func main() {
	helloHandler := HelloHandler{}

	http.Handle("/hello-1", &helloHandler)
	http.HandleFunc("/hello-2", hello)

	http.HandleFunc("/home", home)

	http.HandleFunc("/register", register)
	http.HandleFunc("/success", success)

	fileHandler := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fileHandler))
	// removes /static prefix from the URL path before passing it to the file server,
	// ensuring that the files are correctly served from the static directory.

	fmt.Println("Server listening on port 4000")
	http.ListenAndServe(":4000", nil)
}
