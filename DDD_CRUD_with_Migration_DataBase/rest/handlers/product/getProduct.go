package product

import (
	"ecomerce/util"
	"net/http"
	"strconv"
)

func(h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {

	productId := r.PathValue("id")

	pId,err := strconv.Atoi(productId) 

	if err != nil {

		util.SendError(w,http.StatusBadRequest,"Invalid req body")

		return

	}

	 product,err := h.svc.Get(pId)

	 if err != nil {
		util.SendError(w,http.StatusInternalServerError,"Internal Server Error")
	}

	 if product == nil {
		util.SendError(w,http.StatusNotFound,"Product Not Found")
		return
	 }

	util.SendData(w,http.StatusOK,product)

}