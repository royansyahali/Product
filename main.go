package main

import (
	"erajaya-project/cache"
	"erajaya-project/config"
	"erajaya-project/domain"
	"erajaya-project/repository"
	"erajaya-project/usecase"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var productRepository domain.ProductRepository
var productUseCase domain.ProductUseCase
var productCache cache.ProductCache

func getProducts(ctx *gin.Context) {
	orderBy := ctx.Query("order_by")
	order := ctx.Query("order")

	fmt.Println(orderBy)
	products, err := productUseCase.Get(orderBy, order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, products)
}

func createProduct(ctx *gin.Context) {
	validRequest := true
	validate := validator.New()

	var product domain.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		validRequest = false
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if err := validate.Struct(product); err != nil {
		validRequest = false
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	createdProduct, err := productUseCase.Create(product)
	if err != nil {
		validRequest = false
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	if validRequest {
		ctx.JSON(http.StatusOK, createdProduct)
	}
}

func main() {
	db, _ := config.InitializeDatabase()

	productRepository = repository.NewProductRepository(db)
	productCache = cache.NewRedisCache("localhost:6379", 0, 5)
	productUseCase = usecase.NewProductUseCase(productRepository, productCache)

	router := gin.Default()
	router.GET("/product", getProducts)
	router.POST("/product", createProduct)
	router.Run()

	log.Fatal(http.ListenAndServe(":8080", router))
}
