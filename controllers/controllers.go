package controllers

import (
	"github.com/gin-gonic/gin"
)

func HashPassword(password string) string {
	return ""
}

func VerifyPassword(userPassword string, givenPassword string) (bool, string) {
	return true, ""
}

func Signup() gin.HandlerFunc{
	
}

func Login() gin.HandlerFunc {

}

func ProductViewerAdmin() gin.HandlerFunc {

}

func SearchProduct() gin.HandlerFunc {

}

func SearchProductByQuery() gin.HandlerFunc {

}