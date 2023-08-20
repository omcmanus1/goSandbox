package webserver

import (
	"fmt"
	"net/http"

	"github.com/omcmanus1/goSandbox/internal/types"
)

func CreateServer(port string) error {
	http.Handle("/", &types.CustomHandler{Message: "Welcome to the server"})
	http.Handle("/categories", &types.CustomHandler{Message: "This is the categories page"})
	http.Handle("/secrets", &types.CustomHandler{Message: "Super secret passwords here. Don't look."})
	http.Handle("/blah", &types.CustomHandler{Message: "Blahblahblahblah"})

	fmt.Println("Listening on port ", port)
	return http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
