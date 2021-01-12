package router

import (
	"github.com/labstack/echo/v4"
	"timewise/handler"
	"timewise/middleware"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
	CateHandler handler.CateHandler
	ProductHandler handler.ProductHandler
	OrderHandler handler.OrderHandler
}

func (api *API) SetupRouter() {
	// user
	user := api.Echo.Group("/user")
	user.POST("/sign-up", api.UserHandler.HandleSignUp)
	user.POST("/sign-in", api.UserHandler.HandleSignIn)
	user.GET("/profile", api.UserHandler.HandleProfile, middleware.JWTMiddleware())
	user.GET("/list", api.UserHandler.HandleListUsers, middleware.JWTMiddleware())

	// categories
	categories := api.Echo.Group("/cate",
		middleware.JWTMiddleware())

	categories.POST("/add", api.CateHandler.HandleAddCate)
	categories.PUT("/edit", api.CateHandler.HandleEditCate)
	categories.GET("/detail/:id", api.CateHandler.HandleCateDetail)
	categories.DELETE("/delete/:id", api.CateHandler.HandleDeleteCate)
	categories.GET("/list", api.CateHandler.HandleCateList)

	// products
	product := api.Echo.Group("/product",
		middleware.JWTMiddleware())

	product.POST("/add", api.ProductHandler.HandleAddProduct)
	product.GET("/detail/:id", api.ProductHandler.HandleProductDetail)
	product.GET("/search/:name", api.ProductHandler.HandleSearchProduct)
	product.GET("/cate/:cate", api.ProductHandler.HandleSelectProductByCate)
	product.DELETE("/delete/:id", api.ProductHandler.HandleDeleteProduct)
	product.GET("/list", api.ProductHandler.HandleProductList)
	product.PUT("/edit", api.ProductHandler.HandleEditProduct)

	//order
	order := api.Echo.Group("/order",middleware.JWTMiddleware())

	order.POST("/add", api.OrderHandler.AddToCard)
	order.POST("/confirm", api.OrderHandler.Confirm)
	order.POST("/update", api.OrderHandler.Update)
	order.DELETE("/delete", api.OrderHandler.Remove)
	order.PUT("/edit",api.OrderHandler.UpdateOrderAdmin)
	order.GET("/count", api.OrderHandler.OrderCountItem)
	order.GET("/detail", api.OrderHandler.OrderDetails)
	order.GET("/list", api.OrderHandler.OrderList)
	order.GET("/user/list", api.OrderHandler.OrderByUserId)

}

func (api *API) SetupAdminRouter() {
	//admin
	admin := api.Echo.Group("/admin")
	admin.GET("/token", api.UserHandler.GenToken)
	admin.POST("/sign-in", api.UserHandler.HandleAdminSignIn)
	admin.POST("/sign-up", api.UserHandler.HandleAdminSignUp, middleware.JWTMiddleware())
}