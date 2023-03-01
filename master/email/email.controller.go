package email

import (
	// "encoding/json"
	// "fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gmail-clone.wisnu.net/master/category"
	"gmail-clone.wisnu.net/master/user"
)

func (email *Email) CreateEmail(c *gin.Context) {
	user:= user.User{}
	category := category.Category{}

	err := c.BindJSON(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Bad Request",
		})
		return
	}
	email.ID = 0

	// Get User Id
	resUser := user.CheckUserByEmail(email.From)
	if resUser == nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code":    http.StatusNoContent,
			"message": "User Not Found",
		})
		return
	}

	// Get Category Id
	resCategory,ca := category.CheckIsExistByName("ALL EMAIL",int(resUser.ID))
	if resCategory == 0 {
		c.JSON(http.StatusNoContent, gin.H{
			"code":    http.StatusNoContent,
			"message": "Category Not Found",
		})
		return
	}

	email.CategoryId = int(ca.ID)
	err = email.Create(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Something Went wwong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": email,
	})
}

func (email *Email) GetEmail(c *gin.Context) {
	data := make(map[string]interface{})
	fromQuery := c.Query("from")
	toQuery := c.Query("to")
	pageQuery := c.Query("page")
	limitQuery := c.Query("limit")
	isStarQuery := c.Query("isStar")
	categoryQuery := c.Query("category")

	if fromQuery != "" {
		data["from"] = fromQuery
	}

	if toQuery != "" {
		data["to"] = toQuery
	}

	if categoryQuery != "" {
		categoryId, err := strconv.Atoi(categoryQuery)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":  http.StatusBadRequest,
				"message": "Bad Request",
			})
			return
		}
		data["category_id"] = categoryId
		
	}

	if isStarQuery != ""{
		isStar, err := strconv.ParseBool(isStarQuery)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":  http.StatusBadRequest,
				"message": "Bad Request",
			})
			return
		}
		data["is_star"] = isStar
	}

	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		page = 0
	}

	limit, err := strconv.Atoi(limitQuery)
	if err != nil {
		limit = 15
	}

	res := email.Get(data, page, limit)
	if res == nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code":    http.StatusNoContent,
			"message": "Failed to get data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": res,
	})
}

func (email *Email) UpdateIsStarEmail(c *gin.Context) {
	// Get params
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Wrong format id",
		})
		return
	}

	// Check email is exist by id
	err = email.isExistById(id)
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code":    http.StatusNoContent,
			"message": "Email not found",
		})
		return
	}

	// Get Body
	err = c.BindJSON(&email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Wrong format body",
		})
		return
	}

	// update status is star email
	err = email.UpdateStar(id, email.IsStar)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	// Send json success
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": email,
	})
}

func (email *Email) UpdateCategoryEmail(c *gin.Context){
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Wrong format id",
		})
		return
	}
	
	// Bind Json
	err = c.BindJSON(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Wrong Json",
		})
		return
	}

	// Check Email is Exists
	err = email.isExistById(id)
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code":    http.StatusNoContent,
			"message": "Email not found",
		})
		return
	}

	// Check Category is Exist
	category := category.Category{}
	err = category.CheckIsExistById(email.CategoryId)
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code":    http.StatusNoContent,
			"message": "Category not found",
		})
		return
	}

	// Update Category Email
	err = email.UpdateCategory(id,email.CategoryId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	// Send Json
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"message": email,
	})
}

func (email *Email) DeleteEmail(c *gin.Context){
	// Get Param
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Wrong format id",
		})
		return
	}
	// Check Email Exist
	err = email.isExistById(id)
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code":    http.StatusNoContent,
			"message": "Email not found",
		})
		return
	}

	// // Delete Email
	err = email.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	// Reponse Json
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": err,
	})
}

func (email *Email) GetEmailUser(c *gin.Context){
	data := make(map[string]interface{})
	query := c.Query("query")
	emailQuery := c.Query("email")
	categoryQuery := c.Query("category")
	pageQuery := c.Query("page")
	limitQuery := c.Query("limit")

	if query != "" {
		data["query"] = query
	}

	if emailQuery != "" {
		data["email"] = emailQuery
	}

	if categoryQuery != ""{
		categoryId, err := strconv.Atoi(categoryQuery)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "Bad Request",
			})
			return
		}
		data["category_id"] = categoryId
	}

	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		page = 0
	}

	limit, err := strconv.Atoi(limitQuery)
	if err != nil {
		limit = 15
	}

	
	res := email.GetEmailNew(data,limit,page)
	if res == nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code":    http.StatusNoContent,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": res,
	})

}
