package product

import (
	"ecomerce/database"
	"ecomerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)



func (h *Handler) UpdateProduct (w http.ResponseWriter,r *http.Request) {

		productId := r.PathValue("id")

	pId,err := strconv.Atoi(productId) 

	if err != nil {

		http.Error(w," Please give me a valid product id",400)

		return

	}

		var newProduct database.Products

	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&newProduct) 

	if err != nil {

		fmt.Println(err)

		http.Error(w,"Please Give me a valid json",400)

		return

	}

	newProduct.ID = pId

	database.Update(newProduct)

	util.SendData(w,"Successfully Update Product",200)

}

