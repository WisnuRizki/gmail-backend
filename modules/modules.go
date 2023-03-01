package modules

import (
	"time"
)


type User struct {
	ID        	uint      		`json:"id"`
	FirstName 	string    		`json:"first_name"`
	LastName  	string    		`json:"last_name"`
	Email     	string    		`json:"email"`
	Password  	string    		`json:"password"`
	Category 	[]Category		`gorm:"foreignKey:UserId"`
	CreatedAt 	time.Time 		`json:"created_at"`
	UpdatedAt 	time.Time 		`json:"updated_at"`
}

type Email struct {
	ID        	uint      		`json:"id"`
	From 		string    		`json:"from"`
	To  		string    		`json:"to"`
	Subject     string    		`json:"subject"`
	Message  	string    		`json:"message"`
	Type  		string    		`json:"type"`
	IsStar  	bool    		`json:"isStar"`
	CategoryId	int				`json:"categoryId"`
	Category	*Category		`gorm:"foreignKey:CategoryId"`
	CreatedAt 	time.Time 		`json:"created_at"`
	UpdatedAt 	time.Time 		`json:"updated_at"`
}

type Category struct {
	ID        	uint      		`json:"id"`
	UserId 		int   			`json:"user_id"`
	Name  		string    		`json:"name"`
	CreatedAt 	time.Time 		`json:"created_at"`
	UpdatedAt 	time.Time 		`json:"updated_at"`
}





