package product

import (
	"ecomerce/database"
	"ecomerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func(h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Products

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newProduct) 

	if err != nil {

		fmt.Println(err)

		http.Error(w,"Please Give me a valid json",400)

		return

	}

	createProduct := database.Store(newProduct,)

	util.SendData(w,createProduct,201)

}

