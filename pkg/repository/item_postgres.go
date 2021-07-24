package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/nekruz08/online-store/models"
	"github.com/jmoiron/sqlx"
)

type ItemPostgres struct {
	db  *sqlx.DB
}


func NewItemPostgres(db *sqlx.DB) *ItemPostgres {
	return &ItemPostgres{db: db}
}

// CreateItem - добавляем товары
func (r *ItemPostgres) CreateItem(item models.Item) error  {
	addItem := fmt.Sprintf("INSERT INTO %s(title, categories_id, price, count) VALUES ($1, $2, $3, $4)", itemsTable)
	_, err := r.db.Exec(addItem, item.Title, item.CategoriesID, item.Price, item.Count)
	if err == sql.ErrNoRows{
		return errors.New("ошибка при заполнении товаров")
	}
	if err != nil {
		return err
	}

	return nil
}
// GetItemsByIdCategory - берет товары по ID категории
func (r *ItemPostgres) GetItemsByIdCategory(id int) ([]models.Item, error){
	var items []models.Item

	getItemsById := fmt.Sprintf("SELECT id, title, categories_id, price, count FROM %s WHERE categories_id = $1; ", itemsTable)
	err := r.db.Select(&items, getItemsById, id)
	if err == sql.ErrNoRows{
		return nil, errors.New("не найдено товары")
	}
	if err != nil {
		return nil, err
	}

	return items, err

}


// GetAllItems - берет все товары
func (r *ItemPostgres) GetAllItems() ([]models.Item, error) {
	var items []models.Item

	getItemsById := fmt.Sprintf("SELECT id, title, categories_id, price, count FROM %s; ", itemsTable)
	err := r.db.Select(&items, getItemsById)
	if err == sql.ErrNoRows {
		return nil, errors.New("товары не найдены")
	}
	if err != nil {
		return nil, err
	}

	return items, err

}

//RemoveItemByID - удаляет продукт по ID
func (r *ItemPostgres) RemoveItemByID(id int) (int, error)  {
	deleteItem := fmt.Sprintf("DELETE FROM %s WHERE id = $1 RETURNING id ", itemsTable)
	err := r.db.QueryRow(deleteItem, id).Scan(&id)
	if err == sql.ErrNoRows{
		return 0, errors.New("товар не найден")
	}
	if err != nil {
		return 0, err
	}


	return id, nil
}

// UpdateItemByID - обновляет товар  по ID
func (r *ItemPostgres) UpdateItemByID(item *models.Item) (*models.Item, error) {

	updateItem := fmt.Sprintf("UPDATE %s SET title = $1 ,categories_id = $2,price = $3, count = $4 where id = $5", itemsTable)
	_, err := r.db.Exec(updateItem, item.Title, item.CategoriesID, item.Price, item.Count, item.Id)
	if err == sql.ErrNoRows{
		return nil, errors.New("Товар не найден")
	}
	if err != nil {
		return nil, err
	}

	return item, err
}