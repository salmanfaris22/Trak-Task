package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_name    *string            `bson:"first_name" validate:"required, min=2, max=50"`
	Last_name     *string            `bson:"last" validate:"required, min=2, max=50"`
	Password      *string            `bson:"Password" validate:"required", min=2, max=16"`
	Email         *string            `bson:"email" validate:"email, required"`
	Phone         *string            `bson:"phone" validate:"required", min=10, max=10"`
	Token         *string            `bson:"token"`
	User_type     *string            `bson:"user_type" validate:"required eq=ADMIN|eq=USER"`
	Refresh_token *string            `bson:"refersh_token"`
	Created_at    time.Time          `bson:"created_at"`
	Updated_at    time.Time          `bson:"updated_at"`
	User_id       *string            `bson:"user_id"`
}
