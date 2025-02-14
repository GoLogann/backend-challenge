package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Price       float64            `bson:"price"`
	CategoryID  primitive.ObjectID `bson:"category_id"`
	OwnerID     primitive.ObjectID `bson:"owner_id"`
}
