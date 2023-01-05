package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Marif226/rest-photo-collection/pkg/router"
)

const portNumber = ":8090"

func main() {
	server := &http.Server {
		Addr: portNumber,
		Handler: routes(),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func routes() http.Handler {
	mux := router.NewPathResolver()

	mux.Add("GET /venues/lox/", hello)
	mux.Add("GET /venues/:id/", hello)

	// mux.Add("GET /", index)
	// mux.Add("GET /venues", getVenues)
	// mux.Add("GET /venues/:id", getVenueById)
	// mux.Add("GET /venues/:id/photos/:id", getPhotos)

	// mux.Add("POST /users", createUser)
	// mux.Add("POST /venues", createVenue)
	// mux.Add("POST /venues/:id/photos", createPhoto)

	// mux.Add("PUT /users/:id", updateUser)
	// mux.Add("PUT /venues/:id", update)
	// mux.Add("PUT /venues/:id/photo/:id", updatePhoto)

	// mux.Add("DELETE /venues/:id", deleteVenue)
	// mux.Add("DELETE /venues/:id/photos/:id", deletePhoto)

	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}