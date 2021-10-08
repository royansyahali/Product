package cache

import "erajaya-project/domain"

type ProductCache interface {
	Set(string, domain.Product)
	Get(string, string) []domain.Product
}
