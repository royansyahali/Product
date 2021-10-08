package repository

import (
	"erajaya-project/domain"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r productRepository) Create(product domain.Product) (domain.Product, error) {
	newProduct := domain.Product{
		ID:          uuid.New(),
		Name:        product.Name,
		Price:       product.Price,
		Quantity:    product.Quantity,
		Description: product.Description,
	}
	err := r.Find(product)
	if err != nil {
		return domain.Product{}, err
	}
	err = r.db.Create(&newProduct).Error
	if err != nil {
		return domain.Product{}, errors.Wrap(err, "Failed to create product")
	}
	return newProduct, nil
}

func (r productRepository) Get(orderBy string, order string) ([]domain.Product, error) {
	var filter string
	if orderBy != "" {
		filter = orderBy

		if order != "" {
			filter = fmt.Sprintf("%s %s", orderBy, order)
		}
	}

	fmt.Println(filter)

	var products []domain.Product

	err := r.db.Order(filter).Find(&products).Error
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get products")
	}

	return products, nil
}

func (r productRepository) Find(product domain.Product) error {
	// Get first matched record
	result := r.db.Find(&product, "name = ?", product.Name)
	id := result.RowsAffected
	log.Println(id)
	if id != 0 {
		return errors.New("product already exists")
	}

	return nil
}
