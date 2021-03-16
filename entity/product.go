package entity

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

//Product defines a structure for an item in product catalog
type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	IsAvailable bool    `json:"isAvailable"`
}

// ErrNoProduct is used if no product found
var ErrNoProduct = errors.New("no product found")

// GetProducts returns the JSON file content if available else returns an error.
func GetProducts() ([]byte, error) {
	// Read JSON file
	data, err := ioutil.ReadFile("./data/data.json")
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetProduct takes id as input and returns the corresponding product, else it returns ErrNoProduct error.
func GetProduct(id string) (Product, error) {
	// Read JSON file
	data, err := ioutil.ReadFile("./data/data.json")
	if err != nil {
		return Product{}, err
	}
	// read products
	var products []Product
	err = json.Unmarshal(data, &products)
	if err != nil {
		return Product{}, err
	}
	// iterate through product array
	for i := 0; i < len(products); i++ {
		// if we find one product with the given ID
		if products[i].ID == id {
			// return product
			return products[i], nil
		}
	}
	return Product{}, ErrNoProduct
}

// DeleteProduct takes id as input and deletes the corresponding product, else it returns ErrNoProduct error.
func DeleteProduct(id string) error {
	// Read JSON file
	data, err := ioutil.ReadFile("./data/data.json")
	if err != nil {
		return err
	}
	// read products
	var products []Product
	err = json.Unmarshal(data, &products)
	if err != nil {
		return err
	}
	// iterate through product array
	for i := 0; i < len(products); i++ {
		// if we find one product with the given ID
		if products[i].ID == id {
			products = removeElement(products, i)
			// Write Updated JSON file
			updatedData, err := json.Marshal(products)
			if err != nil {
				return err
			}
			err = ioutil.WriteFile("./data/data.json", updatedData, os.ModePerm)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return ErrNoProduct
}

// AddProduct adds an input product to the product list in JSON document.
func AddProduct(product Product) error {
	// Load existing products and append the data to product list
	var products []Product
	data, err := ioutil.ReadFile("./data/data.json")
	if err != nil {
		return err
	}
	// Load our JSON file to memory using array of products
	err = json.Unmarshal(data, &products)
	if err != nil {
		return err
	}
	// Add new Product to our list
	products = append(products, product)

	// Write Updated JSON file
	updatedData, err := json.Marshal(products)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("./data/data.json", updatedData, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// removeElement is used to remove element from product array at given index
func removeElement(arr []Product, index int) []Product {
	ret := make([]Product, 0)
	ret = append(ret, arr[:index]...)
	return append(ret, arr[index+1:]...)
}
