package models

type Item struct {
	Id int `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
	CategoriesID int `json:"categories_id" db:"categories_id"`
	Price int `json:"price" db:"price"`
	Count int `json:"count" db:"count"`
}


