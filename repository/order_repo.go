package repository

import (
	"timewise/model"
	"context"
)

type OrderRepository interface {
	UpdateStateOrder(context context.Context, order model.Order) error
	UpdateQuantityOrder(context context.Context, userId string, orderId string, quantity int, productId string) error
	UpdateOrder(context context.Context, order model.Order) error
	AddToCard(context context.Context, userId string, card model.Card) (model.Cart, error)
	CountShoppingCard(context context.Context, userId string) (model.OrderCount, error)
	ShoppingCard(context context.Context, userId string, orderId string) (model.Order, error)
	OrderDetailCard(context context.Context, userId string, orderId string) (model.Order, error)
	RemoveCard(context context.Context, userId string, productId string, orderId string)  error
	ListOrder(context context.Context) ([]model.Order, error)
	ListOrderByUserId(context context.Context, userId string) ([]model.Order, error)
	ListDeletedOrderByUserId(context context.Context, userId string) ([]model.Order, error)
}
