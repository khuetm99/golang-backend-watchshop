package repository

import (
	"context"
	"timewise/model"
)

type ProductRepo interface {
	SaveProduct(context context.Context, product model.Product) (model.Product, error)
	UpdateProduct(context context.Context, product model.Product) error
	DeleteProduct(context context.Context, product model.Product) error
	SelectProductById(context context.Context, productId string) (model.Product, error)
	SelectProductByCate(context context.Context, cateId string) ([]model.Product, error)
	SelectProductByName(context context.Context, productName string) ([]model.Product, error)
	SelectProducts(context context.Context) ([]model.Product, error)
}
