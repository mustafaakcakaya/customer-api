package mongo

import (
	"CustomerAPI/internal/customer/types"
	"CustomerAPI/pkg/errors"
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type IRepository interface {
	Create(customer *types.Customer) bool
	Update(id uuid.UUID, customer *types.Customer) bool
	Delete(id uuid.UUID) bool
	GetAll() []types.Customer
	GetById(id uuid.UUID) *types.Customer
	Validate(id uuid.UUID) bool
}

type Repository struct {
	mc *mongo.Collection
}

func NewRepository(mc *mongo.Collection) Repository {
	return Repository{mc: mc}
}

func (repo Repository) Create(customer *types.Customer) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	customer.Id = uuid.NewString()

	customer.CreatedAt = time.Now()
	customer.UpdatedAt = time.Now()
	res, err := repo.mc.InsertOne(ctx, customer)

	if err != nil {
		panic(errors.InsertOneFailed)
	}
	return res.InsertedID != ""
}
func (repo Repository) Update(id uuid.UUID, customer *types.Customer) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	customer.UpdatedAt = time.Now()
	updateModel := bson.M{
		"$set": bson.M{
			"name":                customer.Name,
			"email":               customer.Email,
			"updatedAt":           customer.UpdatedAt,
			"address.addressLine": customer.Address.AddressLine,
			"address.city":        customer.Address.City,
			"address.cityCode":    customer.Address.CityCode,
			"address.county":      customer.Address.County,
		},
	}

	res, err := repo.mc.UpdateOne(ctx, bson.M{"_id": id.String()}, updateModel)
	if err != nil {
		panic(errors.UpdateOneFailed)
	}

	if res.MatchedCount == 0 {
		panic(errors.NotFound)
	}

	return true
}
func (repo Repository) Delete(id uuid.UUID) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	res, err := repo.mc.DeleteOne(ctx, bson.M{"_id": id.String()})
	if err != nil {
		panic(errors.DeleteOneFailed)
	}

	if res.DeletedCount == 0 {
		panic(errors.NotFound)
	}
	return true
}
func (repo Repository) GetAll() []types.Customer {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cur, err := repo.mc.Find(ctx, bson.M{})

	if err != nil && err != mongo.ErrNoDocuments {
		panic(errors.FindFailed)
	}

	items := make([]types.Customer, 0)
	if err = cur.All(ctx, &items); err != nil {
		panic(errors.MongoCursorFailed)
	}

	return items
}
func (repo Repository) GetById(id uuid.UUID) *types.Customer {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	f := bson.M{
		"_id": id.String(),
	}
	customer := new(types.Customer)
	if err := repo.mc.FindOne(ctx, f).Decode(&customer); err != nil {
		panic(errors.NotFound)
	}
	return customer
}
func (repo Repository) Validate(id uuid.UUID) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	f := bson.M{
		"_id": id.String(),
	}
	customer := new(types.Customer)
	if err := repo.mc.FindOne(ctx, f).Decode(&customer); err != nil {
		if err == mongo.ErrNoDocuments {
			return false
		}
		panic(errors.NotFound)
	}
	return true
}
