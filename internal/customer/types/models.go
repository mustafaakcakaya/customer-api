package types

import "time"

// Customer model
type (
	Customer struct {
		Id        string    `bson:"_id" json:"id"`
		Name      string    `bson:"name" json:"name"`
		Email     string    `bson:"email" json:"email"`
		Address   *Address  `bson:"address" json:"address"`
		CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
		UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
	}

	Address struct {
		AddressLine string `bson:"addressLine" json:"addressLine"`
		City        string `bson:"city" json:"city"`
		County      string `bson:"county" json:"county"`
		CityCode    *int   `bson:"cityCode" json:"cityCode"`
	}
)
