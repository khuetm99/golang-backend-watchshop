package handler

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"timewise/banana"
	"timewise/model"
	"timewise/repository"
)

type ProductHandler struct {
	ProductRepo repository.ProductRepo
}

func (p ProductHandler) HandleAddProduct(context echo.Context) error {
	productReq := model.Product{}
	if err := context.Bind(&productReq); err != nil {
		log.Error(err.Error())
		return context.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	products, _ := p.ProductRepo.SelectProducts(context.Request().Context())
	fmt.Println(len(products))

	productId, err := uuid.NewUUID()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	productReq.ProductId = productId.String()

	_, err = p.ProductRepo.SaveProduct(context.Request().Context(), productReq)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return context.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       nil,
	})
}

func (p ProductHandler) HandleProductDetail(context echo.Context) error {
	productId := context.Param("id")
	product, err := p.ProductRepo.SelectProductById(context.Request().Context(), productId)
	if err != nil {
		if err == banana.ProductNotFound {
			return context.JSON(http.StatusNotFound, model.Response{
				StatusCode: http.StatusNotFound,
				Message:    err.Error(),
				Data:       nil,
			})
		}

		return context.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return context.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       product,
	})
}

func (p ProductHandler) HandleSearchProduct(context echo.Context) error {
	productName := context.Param("name")
	product, err := p.ProductRepo.SelectProductByName(context.Request().Context(), productName)
	if err != nil {
		if err == banana.ProductNotFound {
			return context.JSON(http.StatusNotFound, model.Response{
				StatusCode: http.StatusNotFound,
				Message:    err.Error(),
				Data:       nil,
			})
		}

		return context.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return context.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       product,
	})
}

func (p ProductHandler) HandleSelectProductByCate(context echo.Context) error {
	cateId := context.Param("cate")

	product, err := p.ProductRepo.SelectProductByCate(context.Request().Context(), cateId)
	if err != nil {
		if err == banana.ProductNotFound {
			return context.JSON(http.StatusNotFound, model.Response{
				StatusCode: http.StatusNotFound,
				Message:    err.Error(),
				Data:       nil,
			})
		}

		return context.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return context.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       product,
	})
}

func (p ProductHandler) HandleEditProduct(context echo.Context) error {
	productReq := model.Product{}
	if err := context.Bind(&productReq); err != nil {
		log.Error(err.Error())
		return context.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	err := p.ProductRepo.UpdateProduct(context.Request().Context(), productReq)
	if err != nil {
		return context.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusOK,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return context.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Cập nhật sản phẩm thành công",
		Data:       nil,
	})
}

func (p ProductHandler) HandleDeleteProduct(context echo.Context) error {
	productId := context.Param("id")
	product := model.Product{
		ProductId: productId,
	}

	err := p.ProductRepo.DeleteProduct(context.Request().Context(), product)
	if err != nil {
		if err == banana.ProductNotFound {
			return context.JSON(http.StatusNotFound, model.Response{
				StatusCode: http.StatusNotFound,
				Message:    err.Error(),
				Data:       nil,
			})
		}

		return context.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return context.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Delete sản phẩm thành công",
		Data:       nil,
	})
}

func (p ProductHandler) HandleProductList(context echo.Context) error {
	products, err := p.ProductRepo.SelectProducts(context.Request().Context())
	if err != nil {
		return context.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return context.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       products,
	})
}
