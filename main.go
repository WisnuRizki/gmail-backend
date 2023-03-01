package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gmail-clone.wisnu.net/database"
	"gmail-clone.wisnu.net/master/category"
	"gmail-clone.wisnu.net/master/email"
	"gmail-clone.wisnu.net/master/user"
)


func main(){
	r := gin.Default()

	database.ConnectDatabase()
	r.Use(CORSMiddleware())
	port := os.Getenv("PORT")
    if port == "" {
        port = "3001"
    }
	fmt.Println("Connection to database Establish")

	userHandler := user.User{}
	emailHandler := email.Email{}
	categoryHandler := category.Category{}

	userRoute := r.Group("/v1/user")
	{
		userRoute.POST("/register",userHandler.Register)
		userRoute.POST("/login-email",userHandler.LoginEmail)
		userRoute.POST("/login-password",userHandler.LoginPassword)
		userRoute.GET("/check",userHandler.Check)
	}

	emailRoute := r.Group("/v1/email")
	{
		emailRoute.POST("/send-email",emailHandler.CreateEmail)
		emailRoute.GET("/get-email",emailHandler.GetEmailUser)
		emailRoute.PUT("/update-star/:id",emailHandler.UpdateIsStarEmail)
		emailRoute.PUT("/update-category/:id",emailHandler.UpdateCategoryEmail)
		emailRoute.DELETE("/delete-email/:id",emailHandler.DeleteEmail)
	}

	categoryRoute := r.Group("/v1/category")
	{
		categoryRoute.POST("/create-category",categoryHandler.CreateCategory)
		categoryRoute.GET("/get-category",categoryHandler.GetByUserId)
	}

	
	r.Run(fmt.Sprintf(":%s", port)) // listen and serve on 0.0.0.0:8080
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Credentials", "true")
        c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT,DELETE")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}