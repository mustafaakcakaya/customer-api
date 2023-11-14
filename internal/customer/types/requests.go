package types

type (
	CustomerUpsertRequest struct {
		Name    string          `json:"name" validate:"required"`
		Email   string          `json:"email" validate:"required"`
		Address *AddressRequest `json:"address" validate:"required"`
	}

	AddressRequest struct {
		AddressLine string `json:"addressLine"`
		City        string `json:"city" validate:"required"`
		County      string `json:"county" validate:"required"`
		CityCode    *int   `json:"cityCode" validate:"required"`
	}

	ProcessResponse struct {
		IsProcessSuccess bool `json:"isProcessSuccess"`
	}

	ValidateResponse struct {
		IsValidated bool `json:"isValidated"`
	}
)

func (r CustomerUpsertRequest) ToCustomer() *Customer {
	address := new(Address)
	address.AddressLine = r.Address.AddressLine
	address.City = r.Address.City
	address.County = r.Address.County
	address.CityCode = r.Address.CityCode

	return &Customer{
		Name:    r.Name,
		Email:   r.Email,
		Address: address,
	}
}
