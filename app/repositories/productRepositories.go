package repositories

import (
	util "crud_product/app"
	"crud_product/app/models"
	"errors"
	"os"
	"strconv"
	"strings"
)

var products map[string]models.Product = map[string]models.Product{}

var filePathProducts string = "app/data/products.csv"

func GetProducts() (map[string]models.Product, error) {
	data, err := os.ReadFile(filePathProducts)
	if err != nil {
		return map[string]models.Product{}, err
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	for _, line := range lines {
		product, err := strToProduct(strings.TrimSpace(line))
		if err != nil {
			return map[string]models.Product{}, err
		}
		products[product.Id] = product

	}

	return products, nil
}

func GetProductById(id string) (models.Product, error) {
	products, err := GetProducts()
	if err != nil {
		return models.Product{}, err
	}

	product, ok := products[id]
	if !ok {
		return models.Product{}, errors.New("product not found")
	}

	return product, nil
}

func UpdateProducts(productUpdated models.Product) error {
	_, err := GetProductById(productUpdated.Id)
	
	if err != nil {
		return err
	}
	products, err := GetProducts()
	if err != nil {
		return err
	}

	products[productUpdated.Id] = productUpdated

	mapStr := mapToString(products)

	if err := util.WriteToFile(filePathProducts, mapStr); err != nil {
		return err
	}

	return nil
}

func DeleteProductById(id string) error {
	products, err := GetProducts()
	if err != nil {
		return err
	}

	_, ok := products[id]
	if !ok {
		return errors.New("product not found")
	}
	delete(products, id)

	mapStr := mapToString(products)

	if err := util.WriteToFile(filePathProducts, mapStr); err != nil {
		return err
	}

	return nil
}

func AddToRepositories(productToRepostiroies models.Product) error {
	str := productToStr(productToRepostiroies)

	if err := util.AppendToFile(filePathProducts, str); err != nil {
		return err
	}

	return nil
}

func productToStr(product models.Product) string {
	var str string

	str += product.Id + "|" + product.Title + "|" + strconv.FormatFloat(product.Price, 'g', -1, 64) + "\n"

	return str
}

func mapToString(map_product map[string]models.Product) string {
	var str string

	for _, v := range map_product {
		str += v.Id + "|" + v.Title + "|" + strconv.FormatFloat(v.Price, 'g', -1, 64) + "\n"
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
