package controller

import (
	"blocklist/internal/infra/database/mongodb"
	model "blocklist/internal/model/entity"
	"blocklist/internal/model/types"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckSpam(c *gin.Context) {
	// Request body
	var req types.CheckSpamRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error":  err.Error(),
			"msg":    "invalid request body",
			"code":   "INVALID_REQUEST_BODY",
			"status": "FAILED",
		})
		return
	}

	// Get the database connection
	db := mongodb.GetCollection("blocklist")

	// Check if the number is in the database
	var number model.PhoneNumber
	filter := bson.M{"number": req.Number}

	err := db.FindOne(c.Request.Context(), filter).Decode(&number)
	var result types.CheckSpamResponse
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Phone number not found in the blocklist
			result.Number = req.Number
			result.Spam = false

			c.JSON(404, gin.H{
				"msg":    "Phone number not found in blocklist",
				"status": "SUCCESS",
				"result": result,
			})
			return
		}

		c.JSON(500, gin.H{
			"msg":    "Internal Server Error",
			"status": "FAILED",
			"code":   "INTERNAL_SERVER_ERROR",
		})
		return
	}

	result.Number = number.Number
	result.Spam = true

	c.JSON(200, gin.H{
		"msg":    "Phone number found in blocklist",
		"status": "SUCCESS",
		"result": result,
	})
}
