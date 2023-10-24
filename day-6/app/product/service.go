package product

import (
	"context"
	"log"
)

type Repository interface {
	Create(ctx context.Context, req Product) (err error)
}

type Service struct {
	repo PostgresGormRepository
}

func NewService(repo PostgresGormRepository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) CreateProduct(ctx context.Context, req Product) (err error) {
	if err := req.Validate(); err != nil {
		log.Println("error: ", err.Error())
	}

	if err = s.repo.Create(ctx, req); err != nil {
		log.Println("error: ", err.Error())
	}
	return err
}
