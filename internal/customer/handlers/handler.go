package handlers

import (
	customerRepository "CustomerAPI/internal/customer/storages/mongo"
	"CustomerAPI/internal/customer/types"
	"CustomerAPI/pkg/errors"
	"CustomerAPI/pkg/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	customerRepo customerRepository.IRepository
}

func NewHandler(e *echo.Echo, customerRepo customerRepository.IRepository) {
	customerHandler := Handler{customerRepo}

	e.GET("/customers", customerHandler.GetAll)
	e.POST("/customers", customerHandler.Create)
	e.PUT("/customers/:id", customerHandler.Update)
	e.DELETE("/customers/:id", customerHandler.Delete)
	e.GET("/customers/validate/:id", customerHandler.Validate)
}

//GetAll Customers
//@Summary Get All Customers
//@Tags Customers
//@Accept json
//@Produce json
//@Success 200 {object} []types.Customer
//@Failure 500 "Internal Error"
//@Router /customers [get]
func (h Handler) GetAll(c echo.Context) error {
	customers := h.customerRepo.GetAll()
	return c.JSON(http.StatusOK, customers)
}

//Create customer
//@Summary Create new customer
//@Tags Customers
//@Accept json
//@Produce json
//@Param types.CustomerUpsertRequest body types.CustomerUpsertRequest true "Customer"
//@Success 200 {object} types.ProcessResponse
//@Failure 404 "Not Found"
//@Failure 500 "Internal Server Error"
//@Router /customers [post]
func (h Handler) Create(c echo.Context) error {
	req := new(types.CustomerUpsertRequest)
	if _, err := utils.ValidateRequest(c, req); err != nil {
		panic(errors.ValidatorError.WrapDesc(err.Error()))
	}

	res := h.customerRepo.Create(req.ToCustomer())
	return c.JSON(http.StatusOK, types.ProcessResponse{IsProcessSuccess: res})
}

//Update customer
//@Summary Update existing customer
//@Tags Customers
//@Accept json
//@Produce json
//@Param id path string true "id"
//@Param types.CustomerUpsertRequest body types.CustomerUpsertRequest true "Customer to update"
//@Success 200 {object} types.ProcessResponse
//@Failure 400 "Bad Request"
//@Failure 404 "Not Found"
//@Failure 500 "Internal Server Error"
//@Router /customers/{id} [put]
func (h Handler) Update(c echo.Context) error {
	uid, _ := uuid.Parse(c.Param("id"))
	req := new(types.CustomerUpsertRequest)
	if _, err := utils.ValidateRequest(c, req); err != nil {
		panic(errors.ValidatorError.WrapDesc(err.Error()))
	}

	res := h.customerRepo.Update(uid, req.ToCustomer())
	return c.JSON(http.StatusOK, types.ProcessResponse{IsProcessSuccess: res})
}

//Delete Customer
//@Summary Delete existing customer
//@Tags Customers
//@Accept json
//@Produce json
//@Param id path string true "id"
//@Success 200 {object} types.ProcessResponse
//@Failure 404 "Not Found"
//@Failure 500 "Internal Error"
//@Router /customers/{id} [delete]
func (h Handler) Delete(c echo.Context) error {
	uid, _ := uuid.Parse(c.Param("id"))

	res := h.customerRepo.Delete(uid)

	return c.JSON(http.StatusOK, types.ProcessResponse{IsProcessSuccess: res})
}

//Validate Customer
//@Summary Validate of existing customer
//@Tags Customers
//@Accept json
//@Produce json
//@Param id path string true "id"
//@Success 200 {object} types.ValidateResponse
//@Failure 404 "Not Found"
//@Failure 500 "Internal Error"
//@Router /customers/validate/{id} [get]
func (h Handler) Validate(c echo.Context) error {
	uid, _ := uuid.Parse(c.Param("id"))

	res := h.customerRepo.Validate(uid)

	return c.JSON(http.StatusOK, types.ValidateResponse{IsValidated: res})
}
