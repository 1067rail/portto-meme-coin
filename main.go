package main

import (
	"fmt"
	"os"

	"github.com/1067rail/portto-meme-coin/database"
	_ "github.com/1067rail/portto-meme-coin/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

var usersMap = make(map[string]string)

var db *gorm.DB

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func setupRouter() *gin.Engine {
	r := gin.Default()

	// c := controller.NewController(db)

	// v1 := r.Group("/api/v1")
	// {
	// 	coins := v1.Group("coins")
	// 	{
	// 		coins.POST("", c.CreateCoin)
	// 		coins.GET(":id", c.GetCoin)
	// 		coins.PATCH(":id", c.UpdateCoin)
	// 		coins.DELETE(":id", c.DeleteCoin)
	// 		coins.POST(":id/poke", c.PokeCoin)
	// 	}
	// }

	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
