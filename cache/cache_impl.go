package cache

import (
	"encoding/json"
	"erajaya-project/domain"
	"log"
	"sort"
	"time"

	"github.com/go-redis/redis"
)

type ProdustRedis struct {
	Host    string
	DB      int
	Expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) ProductCache {
	return &ProdustRedis{
		Host:    host,
		DB:      db,
		Expires: exp,
	}
}

func (p *ProdustRedis) GetClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: p.Host,
		DB:   p.DB,
	})
}

func (p *ProdustRedis) Set(key string, value domain.Product) {
	client := p.GetClient()
	result, err := json.Marshal(value)
	if err != nil {
		log.Print(err)
	}
	err = client.Set(key, result, p.Expires*time.Minute).Err()
	if err != nil {
		log.Print(err)
	}
}

func (p *ProdustRedis) Get(orderBy, order string) []domain.Product {
	keys := []string{}
	products := []domain.Product{}

	client := p.GetClient()
	val, _ := client.Do("KEYS", "*").Result()
	for _, v := range val.([]interface{}) {
		keys = append(keys, v.(string))
	}
	for i := 0; i < len(keys); i++ {
		val, err := client.Get(keys[i]).Result()
		if err != nil {
			log.Println(err)
		}
		product := domain.Product{}
		err = json.Unmarshal([]byte(val), &product)
		if err != nil {
			log.Print(err)
		}
		products = append(products, product)
	}
	sort.SliceStable(products, func(i, j int) bool {
		return Order(products[i], products[j], orderBy, order)
	})
	log.Println(val)
	return products
}

func Order(val1, val2 domain.Product, orderBy, order string) bool {
	if order == "desc" {
		if orderBy == "name" {
			return val1.Name > val2.Name
		} else if orderBy == "price" {
			return val1.Price > val2.Price
		} else if orderBy == "created_at" {
			return val1.CreatedAt.After(val2.CreatedAt)
		}
	} else if order == "asc" {
		if orderBy == "name" {
			return val1.Name < val2.Name
		} else if orderBy == "price" {
			return val1.Price < val2.Price
		} else if orderBy == "created_at" {
			return val1.CreatedAt.Before(val2.CreatedAt)
		}
	}
	return false
}
