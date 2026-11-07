package repositories

import (
	util "crud_product/app"
	"crud_product/app/models"
	"database/sql"
	"errors"
	"os"
	"strconv"
	"strings"

	_ "modernc.org/sqlite"
)

var (
	filePathProducts string = "app/data/products.csv"
	filePathLastId   string = "app/data/lastId.csv"
	pathDB           string = "app/data/products.db"
)

func GetProducts() ([]models.Product, error) {

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
			id    string
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

func GetProductById(product_id string) (models.Product, error) {
	q := `SELECT id, title, price FROM products WHERE id = ?`

	db, err := getConn()

	var (
		id    string
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

func UpdateProducts(productToUpdate models.Product) (models.Product, error) {
	q := `UPDATE products SET title = ?, price = ? WHERE id = ?`

	db, err := getConn()
	if err != nil {
		return models.Product{}, nil
	}

	_, err = GetProductById(productToUpdate.Id)

	if err != nil {
		return models.Product{}, err
	}

	if _, err := db.Exec(q, productToUpdate.Title, productToUpdate.Price, productToUpdate.Id); err != nil {
		return models.Product{}, err
	}

	return productToUpdate, nil
}

func DeleteProductById(id string) error {

	q := `DELETE FROM products WHERE id = ?`

	db, err := getConn()
	if err != nil {
		return err
	}

	// Проверяем есть ли такой продукт, чтобы его удалить
	_, err = GetProductById(id)
	if err != nil {
		return err
	}

	if _, err := db.Exec(q, id); err != nil {
		return err
	}

	return nil
}

func AddToRepositories(title string, price float64) (models.Product, error) {

	q := `INSERT INTO products (id, title, price) VALUES(?, ?, ?)`

	id, err := getNextId()

	db, err := getConn()
	if err != nil {
		return models.Product{}, err
	}

	productToRepositories := models.Product{Id: id, Title: title, Price: price}

	if err := productToRepositories.ValidProduct(); err != nil {
		return models.Product{}, err
	}

	_, err = GetProductById(productToRepositories.Id)

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

func Init() error {
	db, err := getConn()

	if err != nil {
		return err
	}

	q := `CREATE TABLE IF NOT EXISTS products (id TEXT, title TEXT, price REAL)`

	if _, err := db.Exec(q); err != nil {
		return err
	}

	return nil
}

func getNextId() (string, error) {
	data, err := os.ReadFile(filePathLastId)
	if err != nil {
		return "", err
	}

	if string(data) == "" {
		if err := util.WriteToFile(filePathLastId, "1"); err != nil {
			return "", err
		}

		return "1", nil
	}

	idStr := strings.TrimSpace(string(data))

	lastIdInt, err := strconv.Atoi(idStr)
	if err != nil {
		return "", err
	}

	nextIdInt := lastIdInt + 1

	nextIdStr := strconv.Itoa(nextIdInt)

	if err := util.WriteToFile(filePathLastId, nextIdStr); err != nil {
		return "", err
	}

	return nextIdStr, nil
}

func productToStr(product models.Product) string {
	var str string

	str += product.Id + "|" + product.Title + "|" + strconv.FormatFloat(product.Price, 'g', -1, 64) + "\n"

	return str
}

func sliceToString(sliceProduct []models.Product) string {
	var str string = ""

	for _, product := range sliceProduct {
		str += product.Id + "|" + product.Title + "|" + strconv.FormatFloat(product.Price, 'g', -1, 64) + "\n"
	}

	return str
}

func strToProduct(str string) (models.Product, error) {
	products := strings.Split(str, "|")

	id := products[0]
	title := products[1]
	price, err := strconv.ParseFloat(products[2], 64)
	if err != nil {
		return models.Product{}, err
	}

	product := models.Product{Id: id, Title: title, Price: price}

	return product, nil
}
