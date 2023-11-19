package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"init/models"
	"init/pkg/helpers"
	"net/http"
)

func (c *Handler) Order(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		c.CreateOrder(w, r)
	case "GET":
		var values = r.URL.Query()
		if _, ok := values["id"]; ok {
			c.GetByIDOrder(w, r)
		} else {
			c.GetListOrder(w, r)
		}
	case "PUT":
		c.UpdateOrder(w, r)
	case "DELETE":
		c.UpdateProduct(w, r)
	case "PATCH":
		c.ChangeStatus(w,r)
	}
	
}

func (c *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {

	var createOrder models.CreateOrder
	err := json.NewDecoder(r.Body).Decode(&createOrder)
	if err != nil {
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	

	resp, err := c.storage.Order().Create(&createOrder)
	fmt.Println("ku")
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}
	fmt.Println("noku")

	handleResponse(w, http.StatusCreated, resp)
}

func (c *Handler) GetByIDOrder(w http.ResponseWriter, r *http.Request) {

	var id = r.URL.Query().Get("id")
	fmt.Println("id = " ,id)
	if !helpers.IsValidUUID(id) {
		handleResponse(w, http.StatusBadRequest, "id is not uuid")
		return
	}

	resp, err := c.storage.Order().GetByID(&models.OrderPrimaryKey{ID: id})
	if err == sql.ErrNoRows {
		handleResponse(w, http.StatusBadRequest, "no rows in result set")
		return
	}
	fmt.Println("res = " , resp)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, resp)
}

func (c *Handler) GetListOrder(w http.ResponseWriter, r *http.Request) {

	limit, err := getIntegerOrDefaultValue(r.URL.Query().Get("limit"), 10)
	if err != nil {
		handleResponse(w, http.StatusBadRequest, "invalid query limit")
		return
	}

	offset, err := getIntegerOrDefaultValue(r.URL.Query().Get("offset"), 0)
	if err != nil {
		handleResponse(w, http.StatusBadRequest, "invalid query offset")
		return
	}

	search := r.URL.Query().Get("search")
	if err != nil {
		handleResponse(w, http.StatusBadRequest, "invalid query search")
		return
	}
	fmt.Println("Search inside controller:", search)

	resp, err := c.storage.Order().GetList(&models.GetListOrderRequest{
		Limit:  limit,
		Offset: offset,
		Search: search,
	})
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}
	
	// fmt.Println("res: " , resp)``

	handleResponse(w, http.StatusOK, resp)
}

func (c *Handler) UpdateOrder(w http.ResponseWriter, r *http.Request) {

	var updateOrder models.UpdateOrder
	err := json.NewDecoder(r.Body).Decode(&updateOrder)
	if err != nil {
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	rowsAffected, err := c.storage.Order().Update(&updateOrder)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	if rowsAffected == 0 {
		handleResponse(w, http.StatusBadRequest, "no rows affected")
		return
	}

	resp, err := c.storage.Order().GetByID(&models.OrderPrimaryKey{ID: updateOrder.OrderID})
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusAccepted, resp)
}

func (c *Handler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	var id = r.URL.Query().Get("id")

	if !helpers.IsValidUUID(id) {
		handleResponse(w, http.StatusBadRequest, "id is not uuid")
		return
	}

	err := c.storage.Order().Delete(&models.OrderPrimaryKey{ID: id})
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusNoContent, nil)
}


func (c *Handler) ChangeStatus(w http.ResponseWriter, r *http.Request){
	var updateStatus models.ChangeOrderStatus
	err := json.NewDecoder(r.Body).Decode(&updateStatus)
	if err != nil {
		handleResponse(w, http.StatusBadRequest, err)
		return
	}
	rowsAffected, err := c.storage.Order().ChangeStatus(&updateStatus)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}
	if rowsAffected == 0 {
		handleResponse(w, http.StatusBadRequest, "no rows affected")
		return
	}

	resp, err := c.storage.Order().GetByID(&models.OrderPrimaryKey{ID: updateStatus.ID})
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusAccepted, resp)

}
