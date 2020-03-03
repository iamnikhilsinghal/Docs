package user

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	Db "api/database"
	Models "api/models"

	"gopkg.in/go-playground/validator.v9"
)

func Register(c *gin.Context) {
	client := Db.Connect()
	Reqdata := Models.User{}
	c.BindJSON(&Reqdata)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	col := client.Database("QSS").Collection("users")
	fmt.Println("data for insertion:", Reqdata)

	//Check if email id repeated
	filter := bson.M{"email": bson.M{"$eq": Reqdata.Email}}
	count, err := col.CountDocuments(ctx, filter)
	if count > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "Email Id already registered",
		})
		return
	}

	//Check all fields must be filled
	validate := validator.New()
	err = validate.Struct(&Reqdata)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "Kindly fill all the fields properly",
		})
		return
	}

	Reqdata.Status = "active"
	_, err = col.InsertOne(ctx, Reqdata)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "Failed to insert data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Data inserted successfully",
	})

}
