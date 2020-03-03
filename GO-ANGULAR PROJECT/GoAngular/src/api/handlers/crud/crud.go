package crud

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	Db "api/database"
	Models "api/models"
)

func Read(c *gin.Context) {
	client := Db.Connect()
	var Reqdata []*Models.User

	col := client.Database("QSS").Collection("users")
	filter := bson.M{"status": bson.M{"$eq": "active"}}
	cur, err := col.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var oneDoc Models.User
		err = cur.Decode(&oneDoc)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		Reqdata = append(Reqdata, &oneDoc)
	}
	c.JSON(200, gin.H{
		"code":    http.StatusOK,
		"message": "Data Fetched Successfully",
		"data":    Reqdata,
	})
}

func Update(c *gin.Context) {
	email := c.Param("emailID")
	client := Db.Connect()
	form := Models.User{}
	c.BindJSON(&form)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	col := client.Database("QSS").Collection("users")
	filter := bson.M{"email": email}
	log.Println(form)
	log.Println("email is", email)
	form.Status = "active"
	update := bson.M{"$set": form}
	res, _ := col.UpdateOne(ctx, filter, update)
	if res.ModifiedCount == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "NO record found for Updation",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Record Updated Successfully",
		"data":    res,
	})

}

func Delete(c *gin.Context) {
	client := Db.Connect()
	Reqdata := Models.User{}
	c.BindJSON(&Reqdata)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	col := client.Database("QSS").Collection("users")
	filter := bson.M{"email": bson.M{"$eq": Reqdata.Email}}
	update := bson.M{"$set": bson.M{"status": "disable"}}
	res, _ := col.UpdateOne(ctx, filter, update)

	if res.ModifiedCount == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "NO record found for deletion",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Record Deleted Successfully",
		"data":    res,
	})
}

func Readone(c *gin.Context) {
	client := Db.Connect()
	Reqdata := Models.User{}
	c.BindJSON(&Reqdata)
	var Resdata []*Models.User

	col := client.Database("QSS").Collection("users")
	filter := bson.M{"email": bson.M{"$eq": Reqdata.Email}}
	cur, err := col.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var oneDoc Models.User
		err = cur.Decode(&oneDoc)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		Resdata = append(Resdata, &oneDoc)
	}
	fmt.Println(Resdata)
	c.JSON(200, gin.H{
		"code":    http.StatusOK,
		"message": "Data Fetched Successfully",
		"data":    Resdata,
	})
}

// func Delete(c *gin.Context) {
// 	client := Db.Connect()
// 	Reqdata := Models.User{}
// 	c.BindJSON(&Reqdata)
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	col := client.Database("QSS").Collection("users")
// 	filter := bson.M{"email": bson.M{"$eq": Reqdata.Email}}
// 	res, _ := col.DeleteOne(ctx, filter)
// 	if res.DeletedCount == 0 {
// 		c.JSON(http.StatusOK, gin.H{
// 			"code":    http.StatusOK,
// 			"message": "NO record found for deletion",
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"code":    http.StatusOK,
// 		"message": "Record Deleted Successfully",
// 		"data":    res,
// 	})
// }
