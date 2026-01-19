package handlers

import (
	"ecomerce/database"
	"ecomerce/util"
	"net/http"
	"strconv"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {

	productId := r.PathValue("id")

	pId,err := strconv.Atoi(productId) 

	if err != nil {

		http.Error(w," Please give me a valid product id",400)

		return

	}

	for _,product := range database.ProductsList {

		if product.ID == pId {

			util.SendData(w,product,200)

			return
		}

	}

	util.SendData(w,"No Data Found",404)

}