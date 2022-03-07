package handlers

import (
	"context"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/mosmartin/go-fiber-rest-api/internal/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" validate:"required"`
	Title     string             `json:"title" bson:"title" validate:"required,min=10,max=255"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"created_at,omitempty" validate:"required"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updated_at,omitempty" validate:"required"`
}

type ErrorResponse struct {
	Namespace       string
	StructNamespace string
	FieldField      string
	Tag             string
	Value           string
}

func ValidateProductStruct(p Product) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(p)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &ErrorResponse{
				Namespace:       err.Namespace(),
				StructNamespace: err.StructNamespace(),
				FieldField:      err.Field(),
				Tag:             err.Tag(),
				Value:           err.Param(),
			})
		}
	}

	return errors
}

func CreateProduct(c *fiber.Ctx) error {
	product := Product{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	errors := ValidateProductStruct(product)
	if errors != nil {
		return c.Status(400).JSON(errors)
	}

	client, err := db.GetMongoClient()
	if err != nil {
		return err
	}

	collection := client.Database(os.Getenv("MONGO_DB")).Collection(string(db.ProductsCollection))

	if _, err := collection.InsertOne(context.TODO(), product); err != nil {
		return err
	}

	return c.JSON(product)
}

func GetAllProducts(c *fiber.Ctx) error {
	var products []Product

	// if err := db.Find(&products); err != nil {
	// 	return c.Status(503).JSON(fiber.Map{"message": err.Error()})
	// }

	return c.JSON(products)
}
