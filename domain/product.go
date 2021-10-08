package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          uuid.UUID      `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Name        string         `validate:"required" gorm:"not null" json:"name"`
	Price       int            `validate:"required,min=0" gorm:"not null" json:"price"`
	Description string         `validate:"required" gorm:"not null" json:"description"`
	Quantity    int            `validate:"required,min=0" gorm:"not null" json:"quantity"`
}

type ProductUseCase interface {
	Create(product Product) (Product, error)
	Get(orderBy string, order string) ([]Product, error)
}

type ProductRepository interface {
	Create(product Product) (Product, error)
	Get(orderBy string, order string) ([]Product, error)
	Find(product Product) error
}
