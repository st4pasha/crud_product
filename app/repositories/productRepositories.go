package repositories

import (
	util "crud_product/app"
	"crud_product/app/models"
	"errors"
	"os"
	"strconv"
	"strings"
)

var (
	filePathProducts string = "app/data/products.csv"
	filePathLastId string = "app/data/lastId.csv"
)

func GetProducts() ([]models.Product, error) {

	var products []models.Product

	data, err := os.ReadFile(filePathProducts)
	if err != nil {
		return []models.Product{}, err
	}

	if string(data) == "" {
		return []models.Product{}, nil
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	for _, line := range lines {
		product, err := strToProduct(strings.TrimSpace(line))
		if err != nil {
			return []models.Product{}, err
		}

		products = append(products, product)

	}

	return products, nil
}

func GetProductById(id string) (models.Product, error) {
	products, err := GetProducts()

	if err != nil {
		return models.Product{}, err
	}

	for _, product := range products {
		if product.Id == id {
			return product, nil
		}
	}

	return models.Product{}, errors.New("product not found")

}

func UpdateProducts(productUpdated models.Product) error {
	_, err := GetProductById(productUpdated.Id)

	if err != nil {
		return err
	}

	if err := DeleteProductById(productUpdated.Id); err != nil {
		return err
	}

	if err := AddToRepositories(productUpdated); err != nil {
		return err
	}

	products, err := GetProducts()
	if err != nil {
		return err
	}

	sliceStr := sliceToString(products)

	if err := util.WriteToFile(filePathProducts, sliceStr); err != nil {
		return err
	}

	return nil
}

func DeleteProductById(id string) error {
	products, err := GetProducts()
	if err != nil {
		return err
	}

	for i, product := range products {
		if product.Id == id {
			products[i] = products[(len(products) - 1)]

			products = products[:len(products)-1]

			sliceStr := sliceToString(products)

			if err := util.WriteToFile(filePathProducts, sliceStr); err != nil {
				return err
			}

			return nil

		}
	}

	return errors.New("product not found")
}

func AddToRepositories(productToRepostiroies models.Product) error {
	_, err := GetProductById(productToRepostiroies.Id)

	if err == nil {
		return errors.New("product was created")
	}

	str := productToStr(productToRepostiroies)

	if err := util.AppendToFile(filePathProducts, str); err != nil {
		return err
	}

	return nil
}

func GetNextId() (string, error) {
	data, err := os.ReadFile(filePathLastId)
	if err != nil {
		return "", err
	}

	if string(data) == "" {
		if err  := util.WriteToFile(filePathLastId, "1"); err != nil {
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
