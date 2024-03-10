package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	id             primitive.ObjectID `json:"_id" bson:"_id"`
	firstName      *string            `json:"first_name" validate:"required,min=2,max=30"`
	lastName       *string            `json:"last_name" validate:"required,min=2,max=30"`
	password       *string            `json:"password" validate:"required,min=6"`
	email          *string            `json:"email" validate:"email,required"`
	phone          *string            `json:"phone" validate:"required"`
	token          *string            `json:"token"`
	refreshToken   *string            `json:"refresh_token"`
	createdAt      time.Time          `json:"created_at"`
	updatedAt      time.Time          `json:"update_at"`
	userId         string             `json:"user_id"`
	userCart       []ProductUser      `json:"usercart" bson:"usercart"`
	addressDetails []Address          `json:"address" bson:"address"`
	orderStatus    []Order            `json:"orders" bson:"orders"`
}

type Product struct {
	id     primitive.ObjectID `bson:"_id"`
	name   *string            `json:"product_name"`
	price  *uint64            `json:"price"`
	rating *uint8             `json:"rating"`
	image  *string            `json:"image"`
}

type ProductUser struct {
	productId   primitive.ObjectID `bson:"_id"`
	productName *string            `json:"product_name" bson:"product_name"`
	price       int                `json:"price" bson:"price"`
	rating      *uint              `json:"rating" bson:"rating"`
	image       *string            `json:"image" bson:"image"`
}

type Address struct {
	id      primitive.ObjectID `bson:"_id"`
	house   *string            `json:"house_name" bson:"house_name"`
	street  *string            `json:"street_name" bson:"street_name"`
	city    *string            `json:"city_name" bson:"city_name"`
	pincode *string            `json:"pin_code" bson:"pin_code"`
}

type Order struct {
	id            primitive.ObjectID `bson:"_id"`
	orderCart     []ProductUser      `json:"order_list" bson:"order_list"`
	orderedAt     time.Time          `json:"ordered_at" bson:"ordered_at"`
	price         int                `json:"price" bson:"price"`
	discount      *int               `json:"discount" bson:"discount"`
	paymentMethod Payment            `json:"payment_method" bson:"payment_method"`
}

type Payment struct {
	digital bool
	cod     bool
}
