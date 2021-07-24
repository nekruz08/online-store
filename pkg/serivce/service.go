package service

import (
	"github.com/nekruz08/online-store/models"
	"github.com/nekruz08/online-store/pkg/repository"
)



type Authorization interface {
	CreateUser(user models.User) (int, error)
	// возвращает сгененированный токен и ошибку
	GenerateToken(username, password string) (string, error)
	// он принимает токен и возврашает ID юзера
	ParseToken(token string) (int, error)
	AdminChecker(userId int) (bool, error)

}



type Carts interface {
	CreateCarts(userId int, count int, itemID int) error
	GetCartById(userId int) ([]models.UsersCarts, int,error)
	UpdateCartByID(userID int, idItem int, count int) error
}



type Item interface {
	CreateItem(item models.Item) error
	GetItemsByIdCategory(id int) ([]models.Item, error)
	GetAllItems() ([]models.Item, error)
	RemoveItemByID(id int) (int, error)
	UpdateItemByID(item *models.Item) (*models.Item, error)

}

type Order interface {
	CreateOrder(phone, address string, userId int) error
	GetActiveOrders() ([]*models.Orders, error)
	GetDeliveredOrders()([]*models.Orders, error)
	GetOrderByID(id int)([]*models.Orders, error)
	ConfirmOrder(cartID int) error
}

type Service struct {
	Authorization
	Item
	Carts
	Order
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Item:          NewItemService(repos.Item),
		Carts:         NewCartsService(repos.Carts),
		Order:         NewOrderService(repos.Order),
	}
}