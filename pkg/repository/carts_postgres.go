package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/nekruz08/online-store/models"

	"github.com/jmoiron/sqlx"
	"log"
)

type CartPostgres struct {
	db *sqlx.DB
}

func NewCartPostgres(db *sqlx.DB) *CartPostgres {
	return &CartPostgres{db: db}
}


// CreateCarts - Создает корзину если ее не существует
func (r *CartPostgres) CreateCarts(userId int, count int, itemID int) error {
	var cartId int    // берет ID корзины
	var ItemCount int // таблица для items
	var cartCount int // кол-во которое есть в carts_items


	query := fmt.Sprintf("SELECT id FROM %s WHERE user_id =$1 AND removed = FALSE and delivered = false", cartsTable)
	queryRow := r.db.QueryRow(query, userId).Scan(&cartId)
	if queryRow == sql.ErrNoRows {
		log.Println("Создание корзины")
		createCart := fmt.Sprintf("INSERT INTO %s(user_id) VALUES ($1) returning id", cartsTable)
		err := r.db.QueryRow(createCart, userId).Scan(&cartId)
		if err != nil {
			return errors.New("ошибка при создании корзины")
		}
	}
	checkingCount := fmt.Sprintf("SELECT count FROM %s WHERE id = $1", itemsTable)
	// возвращаем общее количество товаров
	err := r.db.QueryRow(checkingCount, itemID).Scan(&ItemCount)
	if err != nil {
		log.Println("ошибка тут")
		return err
	}
	// делаем проверку чтобы клиент не заказывал больше чем имеется
	if count > ItemCount {
		return errors.New("Недостаточно товаров, осталось: " + string(rune(ItemCount)))
	}
	// проверка на сущ товара в корзине
	newCount := fmt.Sprintf("SELECT count FROM %s WHERE item_id = $1 and  cart_id = $2", cartsItemsTable)
	queryRow2 := r.db.QueryRow(newCount, itemID, cartId).Scan(&cartCount)

	if queryRow2 == sql.ErrNoRows {
		// добавление товаров в корзину
		addItems := fmt.Sprintf("INSERT INTO %s (cart_id, item_id, count) VALUES ($1, $2, $3)", cartsItemsTable)
		_, err = r.db.Exec(addItems, cartId, itemID, 0)
		if err != nil {
			return errors.New("Ошибка при добавлении в корзину")
		}
	}
	updateCount := fmt.Sprintf("UPDATE %s SET count = count + $1 where cart_id = $2 and item_id = $3", cartsItemsTable)
	_, err = r.db.Exec(updateCount, count, cartId, itemID)
	if err != nil {
		log.Println("ошибка при обновлении корзины")
		return errors.New("ошибка при обновлении корзины")
	}
	return nil
}

// GetCartById - смотрит товары в корзине по ID
func (r *CartPostgres) GetCartById(userId int) ([]models.UsersCarts, int, error) {
	var cartsDataById []models.UsersCarts
	var total int

	getCartsByID := fmt.Sprintf("SELECT i.title, ct.count, i.price FROM %s i INNER JOIN %s c ON c.user_id = $1 INNER JOIN %s ct ON i.id = ct.item_id and ct.cart_id = c.id", itemsTable, cartsTable, cartsItemsTable)
	err := r.db.Select(&cartsDataById, getCartsByID, userId)

	for _, value := range cartsDataById {
		total += value.Count * value.Price
	}

	return cartsDataById, total, err
}

// UpdateCartByID - меняет товары в своей корзине
func (r *CartPostgres) UpdateCartByID(userID int, idItem int, count int) error {
	var cartID int
	var itemCount int

	queryRow := fmt.Sprintf("SELECT id FROM %s where user_id = $1", cartsTable)

	err := r.db.QueryRow(queryRow, userID).Scan(&cartID)
	if err == sql.ErrNoRows {
		return errors.New("корзина не найдена")
	}

	getCount := fmt.Sprintf("SELECT count FROM %s WHERE id = $1", itemsTable)
	err = r.db.QueryRow(getCount, idItem).Scan(&itemCount)
	if err == sql.ErrNoRows {
		return errors.New("товар не найден")

	}
	if err != nil {
		return err
	}
	if count > itemCount {
		return errors.New("недастаточно количеств")
	}

	updateItem := fmt.Sprintf("UPDATE %s SET item_id = $1, count = $2", cartsItemsTable)
	_, err = r.db.Exec(updateItem, idItem, count)
	if err != nil {
		return err
	}

	return nil
}
