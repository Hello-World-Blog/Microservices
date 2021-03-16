package main

import (
	"fmt"
	"net/http"

	"github.com/HelloWorld/goProductAPI/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Create new Router
	router := mux.NewRouter()

	// route properly to respective handlers
	router.Handle("/products", handlers.GetProductsHandler()).Methods("GET")
	router.Handle("/products", handlers.CreateProductHandler()).Methods("POST")
	router.Handle("/products/{id}", handlers.GetProductHandler()).Methods("GET")
	router.Handle("/products/{id}", handlers.DeleteProductHandler()).Methods("DELETE")
	router.Handle("/products/{id}", handlers.UpdateProductHandler()).Methods("PUT")

	// Create new server and assign the router
	server := http.Server{
		Addr:    ":9090",
		Handler: router,
	}
	fmt.Println("Staring Product Catalog server on Port 9090")
	// Start Server on defined port/host.
	server.ListenAndServe()
}
