package service

import (
	"github.com/nekruz08/online-store/models"
	"github.com/nekruz08/online-store/pkg/repository"
)

type ItemService struct {
	repo repository.Item // ссылка на БД
}



func NewItemService(repo repository.Item) *ItemService {
	return &ItemService{repo: repo}
}



// CreateItem - добавляем товары
func (s *ItemService) CreateItem(item models.Item) error  {
	return s.repo.CreateItem(item)
}


// GetItemsByIdCategory - берет товары по ID категории
func (s *ItemService) GetItemsByIdCategory(id int) ([]models.Item, error)  {
	return s.repo.GetItemsByIdCategory(id)

}
// GetAllItems - берет все товары
func (s *ItemService) GetAllItems() ([]models.Item, error) {
	return s.repo.GetAllItems()
}

//RemoveItemByID - удаляет продукт по ID
func (s *ItemService) RemoveItemByID(id int) (int, error)  {
	return s.repo.RemoveItemByID(id)
}

// UpdateItemByID - обновляет товар  по ID
func (s *ItemService) UpdateItemByID(item *models.Item) (*models.Item, error)  {
	return s.repo.UpdateItemByID(item)
}