package models

import "time"

type Carts struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	Removed   bool      `json:"removed"`
}

type UsersCarts struct {
	Title string `json:"title" db:"title"`
	Count int    `json:"count" db:"count"`
	Price int    `json:"price" db:"price"`
}

type CartItems struct {
	CartId int `json:"cart_id"`
	ItemId int `json:"item_id"`
	Count  int `json:"count"`
}
