package repository

import (
	"github.com/nekruz08/online-store/models"
	"github.com/jmoiron/sqlx"
)


type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(user, password string) (models.User, error)
	AdminChecker(userId int) (bool, error)
}
type Carts interface {
	CreateCarts(userId int, count int, itemID int) error
	GetCartById(userId int) ([]models.UsersCarts, int, error)
	UpdateCartByID(userID int, idItem int, count int) error
}

// здесь будет связь с postgres
type Item interface {
	CreateItem(item models.Item) error
	GetItemsByIdCategory(id int) ([]models.Item, error)
	GetAllItems() ([]models.Item, error)
	RemoveItemByID(id int) (int, error)
	UpdateItemByID(item *models.Item) (*models.Item, error)
}

type Order interface {
	CreateOrder(phone, address string, userId int) error
	GetAll(cartItems []models.CartItems, cartID int) ([]models.CartItems, error)
	GetActiveOrders() ([]*models.Orders, error)
	GetDeliveredOrders()([]*models.Orders, error)
	GetOrderByID(id int)([]*models.Orders, error)
	ConfirmOrder(cartID int) error
}

type Repository struct {
	Authorization
	Item
	Carts
	Order
}

//TODO: поменяй потом на pgx
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Item:          NewItemPostgres(db),
		Carts:         NewCartPostgres(db),
		Order:         NewOrderPostgres(db),
	}
}
