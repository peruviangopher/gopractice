package main

// ORIGINAL DEMO https://github.com/sausheong/gwp
import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	fmt.Println("hi")
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))

	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", err)

	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	_ = server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err == nil {
		_, err = session(w, r)

		publicTmplFiles := []string{
			"templates/layout.html",
			"templates/public.navbar.html",
			"templates/index.html",
		}

		privateTmplFiles := []string{
			"templates/layout.html",
			"templates/private.navbar.html",
			"templates/index.html",
		}

		var templates *template.Template
		if err != nil {
			templates = template.Must(template.ParseFiles(publicTmplFiles...))
		} else {
			templates = template.Must(template.ParseFiles(privateTmplFiles...))
		}

		_ = templates.ExecuteTemplate(w, "layout", threads)
	}
}