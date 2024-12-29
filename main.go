package main

import (
	"fmt"
	"os"

	"github.com/1067rail/portto-meme-coin/controller"
	"github.com/1067rail/portto-meme-coin/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var usersMap = make(map[string]string)

var db *gorm.DB

func setupRouter() *gin.Engine {
	r := gin.Default()

	c := controller.NewController(db)

	v1 := r.Group("/api/v1")
	{
		coins := v1.Group("coins")
		{
			coins.POST("", c.CreateCoin)
			coins.GET(":id", c.GetCoin)
			coins.PATCH(":id", c.UpdateCoin)
			coins.DELETE(":id", c.DeleteCoin)
			coins.POST(":id/poke", c.PokeCoin)
		}
	}

	return r
}

func main() {

	if _db, err := database.Initialize(); err != nil {
		fmt.Printf("Database initialize error: %v\n", err)
		os.Exit(1)
	} else {
		db = _db
	}

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
