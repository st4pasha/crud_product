package repositories

import (
	"crud_product/services/warehouse/cmd/app/models"
	"database/sql"
	"errors"
)

var pathWarehouseDB string = "services/warehouse/cmd/app/data/warehouse.db"

func CheckCountProduct(product models.Product) error {
	
	warehouseCount, err := GetCountByID(product.Id)
	if err != nil {
		return err
	}

	if warehouseCount - product.Count <= 0 {
		return errors.New("Out of stock")
	}

	return nil
}

func UpdateCountProduct(product models.Product) error {
	q := `UPDATE warehouse SET id = ? WHERE id = ?`

	warehouseCount, err := GetCountByID(product.Id)

	if err != nil {
		return err
	}

	db, err := getConn()

	if err != nil {
		return err
	}

	if _, err := db.Exec(q, warehouseCount - product.Count, product.Id); err != nil {
		return err
	}

	return nil
}
func GetCountByID(id int) (int, error) {

	var count int

	q := `SELECT id FROM warehouse WHRE id = ?`

	db, err := getConn()
	if err != nil {
		return 0, err
	}

	if err := db.QueryRow(q, id).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil

}
func getConn() (*sql.DB, error){
	db, err := sql.Open("sqlite", pathWarehouseDB)
	if err != nil {
		return nil, err
	}

	return db, nil
}
