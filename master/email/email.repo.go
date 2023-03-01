package email

import (
	"errors"

	"gmail-clone.wisnu.net/database"
	"gmail-clone.wisnu.net/modules"
);

type Email modules.Email

func (email *Email) Create(e *Email) error {
	result := database.DB.Create(&e)
	if result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}

func (email *Email) Get(
	query map[string]interface{},
	limit int,
	page int,
) *[]Email{
	data := []Email{}
	offset := (page - 1) * limit

	queryWhere := database.DB
	if len(query) != 0 {
		queryWhere = queryWhere.Where(query)
	}

	result := queryWhere.Limit(limit).Offset(offset).Find(&data)
	if result.RowsAffected == 0 {
		return nil
	}

	return &data
}

func (email Email) isExistById(id int) error {
	result := database.DB.Where(&Email{ID: uint(id)}).Find(&email)
	if result.RowsAffected == 0 {
		return errors.New("No data")
	}

	//fmt.Printf("%d",result.RowsAffected )

	return nil
}

func (email *Email) UpdateStar(id int,status bool) error {
	result := database.DB.Model(&Email{}).Where("id = ?", id).Update("IsStar", status)
	if result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}

func (email *Email) UpdateCategory(id int,categoryId int) error {
	result := database.DB.Model(&Email{}).Where("id = ?", id).Update("category_id", categoryId)
	if result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}

func (email *Email) Delete(id int) error {
	result := database.DB.Where(&Email{ID: uint(id)}).Delete(&email)
	if result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}

func (email *Email) GetEmailNew(query map[string]interface{},limit int,
	page int,) *[]Email {
	data := []Email{}
	queryWhere := database.DB
	conditions := make(map[string]interface{})
	offset := (page - 1) * limit
   
	if len(query) != 0 {
		if query["query"].(string) == "from" || query["query"].(string) == "to"{
			conditions[query["query"].(string)] = query["email"].(string)
			queryWhere = queryWhere.Where(conditions)
		} else if query["query"].(string) == "ALL EMAIL" {
			queryWhere = queryWhere.Where(&Email{From: query["email"].(string)}).Or(&Email{To: query["email"].(string)})
		} else if query["query"].(string) == "STAR" {
			queryWhere = queryWhere.Where(&Email{IsStar: true,From: query["email"].(string)}).Or(&Email{IsStar: true,To: query["email"].(string)})
		} else if query["query"].(string) == "CATEGORY" {
			queryWhere = queryWhere.Where(&Email{CategoryId: query["category_id"].(int)})
		}
	}
	
	result := queryWhere.Limit(limit).Offset(offset).Find(&data)
	if result.RowsAffected == 0 {
		return nil
	}

	return &data
	
}