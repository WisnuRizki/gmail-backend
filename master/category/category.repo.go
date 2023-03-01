package category

import (

	"gmail-clone.wisnu.net/database"
	"gmail-clone.wisnu.net/modules"
);

type Category modules.Category

func (category Category) CheckIsExistByName(name string,userId int) (int64,*Category) {
	result := database.DB.Where(&Category{Name: name,UserId: userId}).Find(&category)
	if result.RowsAffected == 0 {
		return result.RowsAffected,nil
	}
	
	return result.RowsAffected,&category
}

func (category Category) CheckIsExistById(id int) error {
	result := database.DB.Where(&Category{ID: uint(id)}).Find(&category)
	if result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}

func (category Category) Create(c Category) error {
	result := database.DB.Create(&c)
	if result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}

func (category *Category) GetCategoryByUserId(userId int) *[]Category {
	data := []Category{}
	result := database.DB.Where(&Category{UserId: userId}).Find(&data)
	if result.RowsAffected == 0 {
		return nil
	}

	return &data
}