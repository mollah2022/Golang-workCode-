package handlers

import (
	"ecomerce/database"
	"ecomerce/util"
	"net/http"
)



func GetProduct (w http.ResponseWriter,r *http.Request) {

	util.SendData(w,database.ProductsList,200)
}

