package main
import (
	"github.com/gin-gonic/gin"
	"crypto-api/routes"

)




func main()  {
	router := gin.Default()
	
	router.GET("/api/rate", routes.GetPrice)
	router.POST("/api/sendEmails", routes.SendEmails)
	router.POST("/api/subscribe", routes.Subscribe)

	router.Run("localhost:5000")

}
