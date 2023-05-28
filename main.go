package main
import (
	"github.com/gin-gonic/gin"
	"crypto-api/routes"

)




func main()  {
	router := gin.Default()
	
	router.GET("/api/rate", routes.GetPrice)
	router.Run("localhost:5000")

	router.Post("/api/subscribe", services.Subscribe)


}
