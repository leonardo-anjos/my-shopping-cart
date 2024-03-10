package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/leonardo-anjos/my-shopping-cart/controllers"
	"github.com/leonardo-anjos/my-shopping-cart/database"
	"github.com/leonardo-anjos/my-shopping-cart/middleware"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	router.ApplicationRoutes(router)
	router.Use(middleware.Authentication())

	router.Get("/addtocart", app.AddToCart())
	router.Get("/removeitem", app.RemoveItem())
	router.Get("/cartchackout", app.BuyFromCart())
	router.Get("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
