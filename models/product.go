package models

import (
	"time"

	"github.com/jmoiron/sqlx"
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
