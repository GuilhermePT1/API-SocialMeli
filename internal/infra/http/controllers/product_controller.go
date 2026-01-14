package controllers

import (
	"net/http"
	"strconv"

	"github.com/GuilhermePT1/api-social-meli/internal/application/services"
	"github.com/GuilhermePT1/api-social-meli/internal/domain/dto"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	Service *services.ProductService
}

func NewProductController(service *services.ProductService) *ProductController {
	return &ProductController{Service: service}
}

func (c *ProductController) CreateProduct(ctx *gin.Context) {
	var req dto.ProductRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := c.Service.Create(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, dto.ProductResponseDTO{
		ID:    product.ID,
		Name:  product.Name,
		Type:  product.Type,
		Brand: product.Brand,
		Color: product.Color,
		Notes: product.Notes,
	})
}

func (c *ProductController) GetProductById(ctx *gin.Context) {
	productIDStr := ctx.Param("product_id")
	if productIDStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "É necessário preencher o id do produto"})
		return
	}

	productID, err := strconv.ParseUint(productIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID do produto inválido"})
		return
	}

	product, err := c.Service.GetById(uint(productID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (c *ProductController) GetAllProducts(ctx *gin.Context) {
	products, err := c.Service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)
}
