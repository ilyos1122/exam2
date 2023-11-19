package controller

import (

	"encoding/json"
	
	"init/models"
	"init/pkg/helpers"
	"net/http"
)

func (c *Handler) OrderProduct(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		c.CreateOrderProduct(w, r)

	case "DELETE":
		c.DeleteOrderProduct(w, r)
	
	}
	
}

func (c *Handler) CreateOrderProduct(w http.ResponseWriter, r *http.Request) {

	var createOrderProduct models.CreateOrderProduct
	err := json.NewDecoder(r.Body).Decode(&createOrderProduct)
	if err != nil {
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	

	err = c.storage.OrderProduct().Create(&createOrderProduct)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusCreated, err)
}

func (c *Handler) DeleteOrderProduct(w http.ResponseWriter, r *http.Request) {
	var id = r.URL.Query().Get("order_product_id")

	if !helpers.IsValidUUID(id) {
		handleResponse(w, http.StatusBadRequest, "id is not uuid")
		return
	}

	err := c.storage.OrderProduct().Delete(&models.OrderProductPrimaryKey{OrderProductId: id})
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusNoContent, nil)
}

