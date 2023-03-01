package category

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
);

func (category Category) CreateCategory(c *gin.Context){
	// Bind Json
	err := c.BindJSON(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"code": http.StatusBadRequest,
			"message": "Error With Json data",
		})
		return;
	}
	// Check Label is Exist
	category.Name = strings.ToUpper(category.Name)
	row,_ := category.CheckIsExistByName(category.Name,category.UserId)
	if row > 0 {
		c.JSON(http.StatusFound,gin.H{
			"code": http.StatusFound,
			"message": row,
		})
		return
	}
	// // Masukkan Label
	err = category.Create(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"code": http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	// Kirim Response
	c.JSON(http.StatusOK,gin.H{
		"code": http.StatusOK,
		"message": row,
	})
}

func (category *Category) GetByUserId(c *gin.Context){
	// Get Query
	userIdQuery := c.Query("userId")
	userId, err := strconv.Atoi(userIdQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"message": "Bad Request",
		})
		return
	}
	
	// Get Data
	res := category.GetCategoryByUserId(userId)
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code":  http.StatusNoContent,
			"message": category,
		})
		return
	}

	// Response Json
	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"message": res,
	})
}