package product

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type PostgresSQLXRepository struct {
	db *sqlx.DB
}

func NewPostgresSQLXRepository(db *sqlx.DB) PostgresSQLXRepository {
	return PostgresSQLXRepository{db: db}
}

func (p PostgresSQLXRepository) CreateProductData(ctx context.Context, model Product) (err error) {
	query := ``

	statement, err := p.db.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(model)
	if err != nil {
		return err
	}
	return
}
