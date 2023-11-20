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
	slug := helpers.Slugify(name)
	id := uuidv7.Generate().String()
	err := db.Get(&result, "insert into products (id, name, slug, description, price) values ($1, $2, $3, $4, $5) returning "+PRODUCT_COLUMNS, id, name, slug, description, price)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
