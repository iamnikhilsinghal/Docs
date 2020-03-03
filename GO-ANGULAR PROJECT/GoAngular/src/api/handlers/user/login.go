package user

import (
	"context"
	"net/http"
	"time"

	Db "api/database"
	Models "api/models"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Login(c *gin.Context) {
	client := Db.Connect()
	Reqdata := Models.User{}
	c.BindJSON(&Reqdata)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	col := client.Database("QSS").Collection("users")
	filter := bson.M{"email": bson.M{"$eq": Reqdata.Email}}
	result := Models.User{}
	col.FindOne(ctx, filter).Decode(&result)
	if Reqdata.Email != result.Email || Reqdata.Password != result.Password {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "invalid credentials",
		})
		return
	}

	res := structs.Map(result)
	res["Password"] = ""
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Login successfully",
		"data":    res,
	})
}
