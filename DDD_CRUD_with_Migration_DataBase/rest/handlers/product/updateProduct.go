package product

import (
	"ecomerce/domain"
	"ecomerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ReqCreateProducts struct {
    ID int               `json:"id"`
	Title string         `json:"title"`
	Description string   `json:"description"`
	Price float64        `json:"price"`
	ImgUrl string        `json:"imgUrl"`
}

func (h *Handler) UpdateProduct (w http.ResponseWriter,r *http.Request) {

		productId := r.PathValue("id")

	pId,err := strconv.Atoi(productId) 

	if err != nil {

		util.SendError(w,http.StatusBadRequest," Invalid product id")

		return

	}

	var req ReqCreateProducts

	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&req) 

	if err != nil {

		fmt.Println(err)

		util.SendError(w,http.StatusBadRequest," Invalid req body")

		return

	}

	_,err = h.svc.Update(domain.Products{
		ID:pId,
		Title:req.Title,
		Description: req.Description,
		Price: req.Price,
		ImgUrl: req.ImgUrl,

	})

	if err!= nil {
		util.SendError(w,http.StatusInternalServerError,"Internal Server Error")
	}

	util.SendData(w,http.StatusOK,"Successfully Update Product")

}

