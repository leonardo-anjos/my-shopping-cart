package controllers

import "errors"

var (
	ErrorCantFindProduct    = errors.New("can't find the product")
	ErrorCantDecodeProducts = errors.New("cant find the product")
	ErrorUserIdIsNotValid   = errors.New("this user is not valid")
	ErrorCantUpdateUser     = errors.New("cannot add this product to the cart")
	ErrorCantRemoveItemCart = errors.New("cannot remove this item from the cart")
	ErrorCantGetItem        = errors.New("was unable to get the item from the cart")
	ErrorCantBuyCartItem    = errors.New("cannot update the purchase")
)

func addProductToCart() {

}

func removeCartItem() {

}

func buyItemFromCart() {

}

func instantBuyer() {

}
