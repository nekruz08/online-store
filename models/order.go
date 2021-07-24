package models

type Order struct {
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	CartId    int    `json:"cart_id"`
	Delivered bool   `json:"delivered"`
}

type Orders struct {
	Id        int    `json:"id" db:"id"` // ID заказа
	Phone     string `json:"phone" db:"phone"`
	Address   string `json:"address" db:"address"`
	CartID    int    `json:"cart_id" db:"cart_id"`     // ID корзины
	Delivered bool   `json:"delivered" db:"delivered"` // Статус заказа
	Username  string `json:"username" db:"username"`
	Title     string `json:"title" db:"title"`
	Price     int    `json:"price" db:"price"`
	Count     int    `json:"count" db:"count"`
}