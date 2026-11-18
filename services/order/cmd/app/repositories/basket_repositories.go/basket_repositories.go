package basket

import "database/sql"

var pathBasket = "app/data/basket.db"


func AddToRepositories(productID, productCount int) error {

	q := `INSERT INTO basket(product_id, count) VALUES(?,?)`
	db, err := getConn()
	if err != nil {
		return err
	}

	if _, err := db.Exec(q, productID, productCount); err != nil {
		return err
	}

	return nil
}

func getConn() (*sql.DB, error) {
	db, err := sql.Open("sqlite", pathBasket)
	if err != nil {
		return nil, err
	}

	return db, nil
}
