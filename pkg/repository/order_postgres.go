package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/nekruz08/online-store/models"
	"github.com/jmoiron/sqlx"
	"log"
)

type OrderPostgres struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}


// CreateOrder - Создает заказ
func (r *OrderPostgres) CreateOrder(phone, address string, userId int) error {
	var cartId int
	var cartItems []models.CartItems
	tx, _ := r.db.Begin()

	CheckingCart := fmt.Sprintf("SELECT id FROM %s WHERE user_id =$1 AND removed = FALSE and delivered = false", cartsTable)
	queryRow := r.db.QueryRow(CheckingCart, userId).Scan(&cartId)
	if queryRow == sql.ErrNoRows {
		return errors.New("корзина не найдена")

	}

	createOrder := fmt.Sprintf("INSERT INTO %s(phone, address, cart_id) VALUES($1, $2, $3)", orderTable)
	_, err := tx.Exec(createOrder, phone, address, cartId)
	if err != nil {
		tx.Rollback()

		return errors.New("ошибка при оформлении заказа")
	}
	// меняем статус в корзине на true
	updateStatus := fmt.Sprintf("UPDATE %s set delivered = true where user_id = $1", cartsTable)
	_, err = tx.Exec(updateStatus, userId)
	if err != nil {
		return errors.New("ошибка при изменение  статуса")
	}

	//берем все товары из корзины
	allCartItems, err := r.GetAll(cartItems, cartId)
	if err != nil {
		return errors.New("Ошибка при взятии данных")
	}

	for _, cartItem := range allCartItems {
		fmt.Println(cartItem)
		updateItem := fmt.Sprintf("UPDATE %s SET count = count - $1 WHERE id = $2", itemsTable)
		_, err := tx.Exec(updateItem, cartItem.Count, cartItem.ItemId)
		if err == sql.ErrNoRows {
			return errors.New("данныые не найдены, проверьте пожалуйста")
		}
		if err != nil {
			tx.Rollback()
			return err
		}

	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}


// GetAll - берем все товары из корзины, которые заказал клиент
func (r *OrderPostgres) GetAll(cartItems []models.CartItems, cartID int) ([]models.CartItems, error) {
	getAllCartItems := fmt.Sprintf("SELECT cart_id, item_id, count FROM %s where  cart_id = $1", cartsItemsTable)
	rows, err := r.db.Query(getAllCartItems, cartID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for rows.Next() {

		var c models.CartItems
		err := rows.Scan(&c.CartId, &c.ItemId, &c.Count)
		if err != nil {
			return nil, err
		}
		cartItems = append(cartItems, c)

	}

	return cartItems, nil
}

// GetActiveOrders - берет все недоставленные заказы
func (r *OrderPostgres) GetActiveOrders()([]*models.Orders, error)  {
	var orders []*models.Orders

	activeOrders := fmt.Sprintf("SELECT o.id, o.phone,o.address, o.delivered, u.username, i.title, i.price, ci.count, ci.cart_id FROM orders o " +
		"INNER JOIN carts_items ci on o.cart_id = ci.cart_id " +
		"INNER JOIN items i on ci.item_id = i.id " +
		"INNER JOIN carts c on ci.cart_id = c.id " +
		"INNER JOIN users u on c.user_id = u.id " +
		"WHERE  o.delivered = false")
	err := r.db.Select(&orders, activeOrders)
	if err == sql.ErrNoRows{
		return nil, errors.New("заказы не найдены")
	}
	if err != nil {
		return nil, err
	}


	return orders, nil
}

// GetDeliveredOrders - берет все доставленные заказы
func (r *OrderPostgres) GetDeliveredOrders()([]*models.Orders, error)  {
	var orders []*models.Orders

	activeOrders := fmt.Sprintf("SELECT o.id, o.phone,o.address, o.delivered, u.username, i.title, i.price, ci.count, ci.cart_id FROM orders o " +
		"INNER JOIN carts_items ci on o.cart_id = ci.cart_id " +
		"INNER JOIN items i on ci.item_id = i.id " +
		"INNER JOIN carts c on ci.cart_id = c.id " +
		"INNER JOIN users u on c.user_id = u.id " +
		"WHERE  o.delivered = true")
	err := r.db.Select(&orders, activeOrders)
	if err == sql.ErrNoRows{
		return nil, errors.New("заказы не найдены")
	}
	if err != nil {
		return nil, err
	}


	return orders, nil
}



//GetOrderByID - клиент может посмотреть, что он заказал
func (r *OrderPostgres) GetOrderByID(id int)([]*models.Orders, error)  {
	log.Println(id, "ID который получили")
	var orders []*models.Orders
	// Нужна ли сумма?
	activeOrders := fmt.Sprintf("SELECT o.id, o.phone,o.address, o.delivered, u.username, i.title, i.price, ci.count, ci.cart_id FROM orders o " +
		"INNER JOIN carts_items ci on o.cart_id = ci.cart_id " +
		"INNER JOIN items i on ci.item_id = i.id " +
		"INNER JOIN carts c on ci.cart_id = c.id " +
		"INNER JOIN users u on c.user_id = u.id " +
		"WHERE  u.id = $1")
	err := r.db.Select(&orders, activeOrders, id)
	if err == sql.ErrNoRows{
		return nil, errors.New("заказы не найдены")
	}
	if err != nil {
		return nil, err
	}


	return orders, nil
}
// ConfirmOrder - потверждение заказа
func (r *OrderPostgres) ConfirmOrder(cartID int) error  {
	orderStatus := fmt.Sprintf("UPDATE %s SET delivered = true WHERE cart_id = $1", orderTable)
	_, err := r.db.Exec(orderStatus, cartID)
	if err != nil {
		return err
	}

	return nil
}