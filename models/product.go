package models

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/moroz/omise/helpers"
	"github.com/moroz/uuidv7-go"
)

type Product struct {
	ID          uuidv7.UUID `db:"id"`
	Name        string      `db:"name"`
	Slug        string      `db:"slug"`
	Description *string     `db:"description"`
	Price       int         `db:"price"`
	InsertedAt  time.Time   `db:"inserted_at"`
	UpdatedAt   time.Time   `db:"updated_at"`
}

const PRODUCT_COLUMNS = "id, name, slug, description, price, inserted_at, updated_at"

func ListProducts(db *sqlx.DB) ([]Product, error) {
	result := []Product{}
	err := db.Select(&result, "select "+PRODUCT_COLUMNS+" from products order by id")
	return result, err
}

func CreateProduct(db *sqlx.DB, name string, price int, description string) (*Product, error) {
	result := Product{}
	slug := helpers.GenerateSlug(name)
	id := uuidv7.Generate().String()
	err := db.Get(&result, "insert into products (id, name, slug, description, price) values ($1, $2, $3, $4, $5) returning "+PRODUCT_COLUMNS, id, name, slug, description, price)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetProductBySlug(db *sqlx.DB, slug string) (*Product, error) {
	result := Product{}
	err := db.Get(&result, "select "+PRODUCT_COLUMNS+" from products where slug = $1", slug)
	if err != nil {
		return nil, err
	}
	return &result, err
}

func GetProductById(db *sqlx.DB, id string) (*Product, error) {
	result := Product{}
	err := db.Get(&result, "select "+PRODUCT_COLUMNS+" from products where id = $1", id)
	if err != nil {
		return nil, err
	}
	return &result, err
}

func UpdateProduct(db *sqlx.DB, product *Product, name string, price int, description string) (*Product, error) {
	result := Product{}
	err := db.Get(&result, "update products set name = $1, price = $2, description = $3, updated_at = now() at time zone 'utc' where id = $4 returning "+PRODUCT_COLUMNS,
		name, price, description, product.ID.String())
	if err != nil {
		return nil, err
	}
	return &result, err
}
