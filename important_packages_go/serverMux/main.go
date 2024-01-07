package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello World!"))
	// })

	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog{title: "My blog"})
	http.ListenAndServe(":8080", mux)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

type blog struct {
	title string
}

// ServeHTTP implements http.Handler.
func (blog) ServeHTTP(http.ResponseWriter, *http.Request) {
	panic("unimplemented")
}

func (b blog) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
