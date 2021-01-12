package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/labstack/gommon/log"
	"github.com/lib/pq"
	"time"
	"timewise/banana"
	"timewise/db"
	"timewise/model"
)

type ProductRepoImpl struct {
	sql *db.SQL
}

// NewUserRepo create object working with user logic
func NewProductRepo(sql *db.SQL) ProductRepo {
	return &ProductRepoImpl{
		sql: sql,
	}
}

func (p ProductRepoImpl) SaveProduct(context context.Context, product model.Product) (model.Product, error) {
	statement := `
		INSERT INTO products(
		  		product_id, product_name, product_image, product_des, quantity, 
		  		sold_items, created_at, updated_at, price, cate_id) 
          VALUES(:product_id, :product_name, :product_image,:product_des, :quantity, 
          		 :sold_items, :created_at, :updated_at, :price, :cate_id)	 
	`

	now := time.Now()
	product.CreatedAt = now
	product.UpdatedAt = now

	_, err := p.sql.Db.NamedExecContext(context, statement, product)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return product, errors.New("Sản phẩm này đã tồn tại")
			}
		}
		return product, errors.New("Tạo Sản phẩm thất bại")
	}

	return product, nil
}


func (p ProductRepoImpl) SelectProductById(context context.Context, productId string) (model.Product, error) {
	var product = model.Product{}

	statement := `SELECT * FROM products WHERE product_id=$1`
	err := p.sql.Db.GetContext(context, &product, statement, productId)

	if err != nil {
		if err == sql.ErrNoRows {
			return product, banana.ProductNotFound
		}
		log.Error(err.Error())
		return product, err
	}

	return product, nil
}

func (p ProductRepoImpl) SelectProductByCate(context context.Context, cateId string) ([]model.Product, error) {
	var products = []model.Product{}

	statement := `SELECT
	      products.*,
	      categories.cate_name
	    FROM products 
	      INNER JOIN categories ON products.cate_id = categories.cate_id AND
		  products.cate_id= $1 ORDER BY product_name`
	err := p.sql.Db.SelectContext(context, &products, statement, cateId)

	if err != nil {
		return products, err
	}

	return products, nil
}


func (p ProductRepoImpl) SelectProductByName(context context.Context, productName string) ([]model.Product, error) {
	products := []model.Product{}

	statement := `SELECT
	      products.*,
	      categories.cate_name
	    FROM products 
	      INNER JOIN categories ON products.cate_id = categories.cate_id AND
		  product_name ILIKE '%' || $1 || '%' ORDER BY product_name`

	err := p.sql.Db.SelectContext(context, &products, statement, productName)

	if err != nil {
		return products, err
	}

	return products, nil
}


func (p ProductRepoImpl) UpdateProduct(context context.Context, product model.Product) error {
	sqlStatement := `
		UPDATE products
		SET 
		    product_name = :product_name,
			product_image = :product_image,
			product_des = :product_des,
		    price = :price,
		    quantity = :quantity,
		    sold_items = :sold_items,
			cate_id = :cate_id,
		    updated_at = :updated_at
		WHERE 
			product_id =:product_id;
`

	product.UpdatedAt = time.Now()

	result, err := p.sql.Db.NamedExecContext(context, sqlStatement, product)
	if err != nil {
		return err
	}

	count, _ := result.RowsAffected()
	if count == 0 {
		return errors.New("Update thất bại")
	}

	return nil
}

func (p *ProductRepoImpl) DeleteProduct(context context.Context, product model.Product) (error) {
	sqlStatement := ` 
		DELETE FROM products
		WHERE product_id = $1;
	`
	// Trước khi xoá nên kiểm tra sản phẩm này có thuộc về user này hay không
	result, err := p.sql.Db.ExecContext(context, sqlStatement, product.ProductId)
	count, _ := result.RowsAffected()
	if count == 0 {
		return errors.New("Delete thất bại")
	}
	return err
}

func (p ProductRepoImpl) SelectProducts(context context.Context) ([]model.Product, error) {
	products := []model.Product{}
	sql := `SELECT
	      products.*,
	      categories.cate_name
	    FROM products 
	      INNER JOIN categories ON products.cate_id = categories.cate_id;`
	err := p.sql.Db.SelectContext(context, &products, sql)
	if err != nil {
		return products, err
	}

	return products, nil
}




