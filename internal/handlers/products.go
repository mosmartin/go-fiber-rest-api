package handlers

import (
	"context"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mosmartin/go-fiber-rest-api/internal/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updated_at,omitempty"`
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
