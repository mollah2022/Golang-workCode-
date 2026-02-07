package product

import (
	"ecomerce/domain"
	"ecomerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqCreateProduct struct {
    
	Title string         `json:"title"`
	Description string   `json:"description"`
	Price float64        `json:"price"`
	ImgUrl string        `json:"imgUrl"`
}

func(h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req ReqCreateProduct

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&req) 

	if err != nil {

		fmt.Println(err)

		util.SendError(w,http.StatusBadRequest,"Invalid Req body")

		return

	}

	createProduct,err := h.svc.Creat(domain.Products{
		Title: req.Title,
		Description: req.Description,
		Price: req.Price,
		ImgUrl: req.ImgUrl,
	})

	if err != nil {
		util.SendError(w,http.StatusBadRequest,"Internal Server Error")
		return
	}

	util.SendData(w,http.StatusCreated,createProduct)

}

