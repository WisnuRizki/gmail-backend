package user

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gmail-clone.wisnu.net/master/category"
);


func (user *User) Register(c *gin.Context){
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"code": http.StatusBadRequest,
			"message": "Error With Json data",
		})
		return;
	}

	result := user.CheckUserByEmail(user.Email)
	if result != nil {
		c.JSON(http.StatusFound,gin.H{
			"code": http.StatusFound,
			"message": "User Already Exists",
		})
		return;
	}

	hashPassword,err := bcrypt.GenerateFromPassword([]byte(user.Password),10)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message": "Something Went Wrong",
		})
		return
	}

	user.Password = string(hashPassword)
	user.ID = 0


	err = user.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"code": http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return;
	}

	// Create Category
	category := category.Category{}
	category.UserId = int(user.ID)
	category.Name = "ALL EMAIL"
	err = category.Create(category)
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code":    http.StatusNoContent,
			"message": "Category not found",
		})
		return
	}


	c.JSON(http.StatusOK,gin.H{
		"code": http.StatusOK,
		"message": "Success Create New User",
	})
}

func (user *User) LoginEmail(c *gin.Context){
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"code": http.StatusBadRequest,
			"message": "Error With Json data",
		})
		return;
	}

	result := user.CheckUserByEmail(user.Email)
	if result == nil {
		c.JSON(http.StatusNoContent,gin.H{
			"code": http.StatusNoContent,
			"message": "User Not Found",
		})
		return;
	}

	b, _ := json.Marshal(&result)
	var response map[string]interface{}
	_ = json.Unmarshal(b, &response)

	delete(response,"password")
	delete(response,"created_at")
	delete(response,"updated_at")

	c.JSON(http.StatusOK,gin.H{
		"code": http.StatusOK,
		"message": "Found User",
		"data": response,
	})
}

func (user *User) LoginPassword(c *gin.Context){
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"code": http.StatusBadRequest,
			"message": "Error With Json data",
		})
		return;
	}

	result := user.CheckUserByEmail(user.Email)
	if result == nil {
		c.JSON(http.StatusNoContent,gin.H{
			"code": http.StatusNoContent,
			"message": "User Not Found",
		})
		return;
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password),[]byte(user.Password))
	if err != nil {
		c.JSON(http.StatusNotAcceptable,gin.H{
			"code": http.StatusNotAcceptable,
			"message": "Wrong Password",
		})
		return;
	}

	b, _ := json.Marshal(&result)
	var response map[string]interface{}
	_ = json.Unmarshal(b, &response)

	delete(response,"password")
	delete(response,"created_at")
	delete(response,"updated_at")

	c.JSON(http.StatusOK,gin.H{
		"code": http.StatusOK,
		"message": "Found User",
		"data": response,
	})
}

func (user *User)Check(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"code": http.StatusOK,
		"message": "Ok",
	})
}