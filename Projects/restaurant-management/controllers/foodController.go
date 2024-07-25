package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func round(num float64) int {
	fmt.Print(num)
	return 0
}

func toFixed(num float64, precision int) float64 {
	return 0
}
