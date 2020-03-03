package main

import (
	Crud "api/handlers/crud"
	User "api/handlers/user"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//Cross-Origin Resource Sharing(CORS)
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max, Origin")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})

	r.POST("/register", User.Register)
	r.POST("/login", User.Login)
	r.GET("/read", Crud.Read)
	r.PUT("/update/:emailID", Crud.Update)
	r.POST("/delete", Crud.Delete)
	r.POST("/readone", Crud.Readone)
	r.Run()
}
