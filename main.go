package main

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
	Image    string `json:"image"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Set headers for CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		// Handle pre-flight request
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Return list of products
		products := []Product{
			{Name: "Coca Cola", Price: 100, Quantity: 2, Image: "https://cdn.shopify.com/s/files/1/0268/0622/2893/products/PNG_0005_1706067-copy.png?v=1617186406"},
			{Name: "Fanta", Price: 200, Quantity: 3, Image: "https://cdn.shopify.com/s/files/1/0558/2687/4437/products/VFT204521-fanta-orange-355ml-1_600x.png?v=1670287797"},
			{Name: "Bringles", Price: 300, Quantity: 5, Image: "https://framestrapimaster.blob.core.windows.net/assets/images/FQAT_1730100_2x_2x_282d7e705b.png"},
		}

		json.NewEncoder(w).Encode(products)
	})

	http.ListenAndServe(":8081", nil)
}
