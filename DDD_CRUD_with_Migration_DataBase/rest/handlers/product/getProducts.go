package product

import (
	"ecomerce/util"
	"net/http"
)



func (h *Handler) GetProducts (w http.ResponseWriter,r *http.Request) {

	productList,err := h.svc.List()

	if err !=nil {
		util.SendError(w,http.StatusInternalServerError,"Internal Server Error")
	}

	util.SendData(w,http.StatusOK,productList)
}

