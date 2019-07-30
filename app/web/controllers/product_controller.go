package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/zidni722/pawoon-product/app/repositories/interface"	
	"github.com/zidni722/pawoon-product/app/models"
	"github.com/zidni722/pawoon-product/app/web/response"
)

type ProductController struct {
	Db                *gorm.DB
	ProductRepository _interface.IProductRepository
}

func NewProductController(db *gorm.DB, productRepository _interface.IProductRepository) *ProductController {
	return &ProductController{
		Db:                db,
		ProductRepository: productRepository,
	}
}

func (c *ProductController) GetProductByOutlet(ctx iris.Context) {
	outletUUID, _ := ctx.Params().Get("outletUuid")

	var products models.Product

	c.ProductRepository.FindByUuid(c.Db, &products, outletUUID)

	if products == (models.Product{}) {
		response.ErrorResponse(ctx, response.UNPROCESSABLE_ENTITY, "Product doesn't exists.")
		return
	}

	productResponse := response.NewProductResponse(c.Db)
	result := productResponse.New(detailItem)

	response.SuccessResponse(ctx, response.OK, response.OK_MESSAGE, result)
}
