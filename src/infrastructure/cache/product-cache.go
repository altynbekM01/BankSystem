package cache

import "GolangwithFrame/src/domain/model"

type ProductCache interface {
	Set(key string, value *model.Product)
	Get(key string) *model.Product
}
