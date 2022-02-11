package main

import (
	"fmt"
	"net/http"
	"os"

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
		Handler: handlers.AuthHandler(router),
	}
	fmt.Println("Staring Product Catalog server on Port 9090")
	// Start Server on defined port/host.
	isAzureDeployment := os.Getenv("AZURE_DEPLOYMENT")
	if isAzureDeployment == "TRUE" {
		fmt.Println("Running on Azure")
		err := server.ListenAndServe()
		if err != nil {
			fmt.Println("Failed to start HTTP Server")
		}
	} else {
		err := server.ListenAndServeTLS("server.crt", "server.key")
		if err != nil {
			fmt.Printf("Failed to start HTTPS server: %s", err.Error())
		}
	}
}
