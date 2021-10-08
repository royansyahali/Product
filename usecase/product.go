package usecase

import (
	"erajaya-project/cache"
	"erajaya-project/domain"
)

type useCase struct {
	productRepository domain.ProductRepository
	productCache      cache.ProductCache
}

func NewProductUseCase(productRepository domain.ProductRepository, productCache cache.ProductCache) domain.ProductUseCase {
	return &useCase{
		productRepository: productRepository,
		productCache:      productCache,
	}
}

func (u *useCase) Create(product domain.Product) (domain.Product, error) {
	createdProduct, err := u.productRepository.Create(product)
	if err != nil {
		return domain.Product{}, err
	}
	return createdProduct, nil
}

func (u *useCase) Get(orderBy string, order string) ([]domain.Product, error) {
	cache := u.productCache.Get(orderBy, order)
	if len(cache) == 0 {
		products, err := u.productRepository.Get(orderBy, order)
		if err != nil {
			return nil, err
		}
		for _, v := range products {
			u.productCache.Set(v.ID.String(), v)
		}
		return products, nil
	}
	return cache, nil

}
