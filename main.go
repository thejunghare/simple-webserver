package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// handler to return 404
	if r.URL.Path != "/hello" {
		// http.Error(w, "404", http.StatusNotFound)
		http.Handle("404", http.NotFoundHandler())
		return
	}

	fmt.Fprint(w, "Hello")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	/* http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		http.ServerFile(w, r, "./static/form.html")
	}) */
	// http.Handle("/form", http.FileServer(http.Dir("./static/form.html")))
	// http.ServeFile(w, r, "./static/form.html")

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err :%v", err)
		return
	}

	fmt.Fprintf(w, "Post request successful\n")

	email := r.FormValue("email")
	password := r.FormValue("password")

	fmt.Fprintf(w, "Email = %s\n", email)
	fmt.Fprintf(w, "Password = %s\n", password)
}

func main() {
	fmt.Println("Starting server at port 8080")

	/* http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hello") // This will just say hello

		// http.FileServer(http.Dir("./static/index.html"))
		http.ServeFile(w, r, "./static/index.html")
	}) */

	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
