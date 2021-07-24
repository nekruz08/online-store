package service

import (
	"github.com/nekruz08/online-store/models"
	"github.com/nekruz08/online-store/pkg/repository"
)

type OrderService struct {
	repo   repository.Order
}


func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{repo: repo}
}


// CreateOrder - Создает заказ
func (s *OrderService) CreateOrder(phone, address string, userId int) error {
	return s.repo.CreateOrder(phone, address, userId)
}


// GetActiveOrders - берет все недоставленные заказы
func (s *OrderService) GetActiveOrders() ([]*models.Orders, error) {
	return s.repo.GetActiveOrders()

}

// GetDeliveredOrders - берет все доставленные заказы
func (s *OrderService) GetDeliveredOrders()([]*models.Orders, error)  {
	return s.repo.GetDeliveredOrders()

}

//GetOrderByID - клиент может посмотреть, что он заказал
func (s *OrderService) GetOrderByID(id int)([]*models.Orders, error)   {
	return s.repo.GetOrderByID(id)

}

// ConfirmOrder - потверждение заказа
func (s *OrderService)  ConfirmOrder(cartID int) error  {
	return s.repo.ConfirmOrder(cartID)

}