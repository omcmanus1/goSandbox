package webserver

import (
	"fmt"
	"net/http"
)

func CreateServer(port string) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the server")
	})
	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is the categories page")
	})
	http.HandleFunc("/secrets", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Super secret passwords here. Don't look.")
	})
	http.HandleFunc("/blah", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Blahblahblahblah")
	})

	fmt.Println("Listening on port ", port)
	return http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
