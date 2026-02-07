package product

import (
	"ecomerce/util"
	"fmt"
	"net/http"
	"strconv"
)



func(h *Handler) DeleteProduct (w http.ResponseWriter,r *http.Request) {

		productId := r.PathValue("id")

	pId,err := strconv.Atoi(productId) 

	if err != nil {

	   fmt.Println(err)

		util.SendError(w,http.StatusBadRequest,"Invalid Product ID")

		return

	}

	err = h.svc.Delete(pId)

	if err != nil {
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}

   util.SendData(w,201,"Successfully deleted Product")

}

