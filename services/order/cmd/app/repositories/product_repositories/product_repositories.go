package product

import (
	"crud_product/services/order/cmd/app/models"
	"database/sql"
	"errors"

	_ "modernc.org/sqlite"
)

var (
	pathDB string = "app/data/products.db"
)

func Get() ([]models.Product, error) {

	var products []models.Product

	q := `SELECT id, title, price FROM products`

	db, err := getConn()
	if err != nil {
		return []models.Product{}, err
	}

	rows, err := db.Query(q)

	if err != nil {
		return []models.Product{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			id    int
			title string
			price float64
		)

		if err := rows.Scan(&id, &title, &price); err != nil {
			return []models.Product{}, err
		}

		product := models.Product{Id: id, Title: title, Price: price}

		products = append(products, product)
	}

	if rows.Err() != nil {
		return []models.Product{}, rows.Err()
	}

	return products, nil
}

func GetById(product_id int) (models.Product, error) {
	q := `SELECT id, title, price FROM products WHERE id = ?`

	db, err := getConn()

	var (
		id    int
		title string
		price float64
	)

	err = db.QueryRow(q, product_id).Scan(&id, &title, &price)

	if err == sql.ErrNoRows {
		return models.Product{}, sql.ErrNoRows
	}

	if err != nil {
		return models.Product{}, err
	}

	product := models.Product{Id: id, Title: title, Price: price}

	return product, nil

}

func Update(productToUpdate models.Product) (models.Product, error) {
	q := `UPDATE products SET title = ?, price = ? WHERE id = ?`

	db, err := getConn()
	if err != nil {
		return models.Product{}, nil
	}

	_, err = GetById(productToUpdate.Id)

	if err != nil {
		return models.Product{}, err
	}

	if _, err := db.Exec(q, productToUpdate.Title, productToUpdate.Price, productToUpdate.Id); err != nil {
		return models.Product{}, err
	}

	return productToUpdate, nil
}

func DeleteById(id int) error {

	q := `DELETE FROM products WHERE id = ?`

	db, err := getConn()
	if err != nil {
		return err
	}

	// Проверяем есть ли такой продукт, чтобы его удалить
	_, err = GetById(id)
	if err != nil {
		return err
	}

	if _, err := db.Exec(q, id); err != nil {
		return err
	}

	return nil
}

func AddToRepositories(title string, price float64) (models.Product, error) {

	var id int
	q := `INSERT INTO products (id, title, price) VALUES(?, ?, ?)`

	products, err := Get()

	if err != nil {
		return models.Product{}, err
	}

	if len(products) == 0 {
		id = 1
	} else {
		id = products[len(products)-1].Id + 1
	}

	db, err := getConn()
	if err != nil {
		return models.Product{}, err
	}

	productToRepositories := models.Product{Id: id, Title: title, Price: price}

	if err := productToRepositories.Valid(); err != nil {
		return models.Product{}, err
	}

	_, err = GetById(productToRepositories.Id)

	if err == nil {
		return models.Product{}, errors.New("product was created")
	}

	if _, err := db.Exec(q, productToRepositories.Id, productToRepositories.Title, productToRepositories.Price); err != nil {
		return models.Product{}, err
	}

	return productToRepositories, nil
}

func getConn() (*sql.DB, error) {
	db, err := sql.Open("sqlite", pathDB)
	if err != nil {
		return nil, err
	}

	return db, nil
}
