package service

import (
	"github.com/nekruz08/online-store/models"
	"github.com/nekruz08/online-store/pkg/repository"
)

type CartsService struct {
	repo repository.Carts
}


func NewCartsService(repo repository.Carts) *CartsService {
	return &CartsService{repo: repo}
}

// CreateCarts - Создает корзину если ее не существует
func (s *CartsService) CreateCarts(userId int, count int, itemID int) error  {

	return s.repo.CreateCarts(userId, count, itemID)

}
// GetCartById - смотрит товары в корзине по ID
func (s *CartsService) GetCartById(userId int) ([]models.UsersCarts, int,error) {
	return s.repo.GetCartById(userId)
}

// UpdateCartByID - меняет товары в своей корзине
func (s *CartsService) UpdateCartByID(userID int, idItem int, count int) error  {
	return s.repo.UpdateCartByID(userID, idItem, count)
}
