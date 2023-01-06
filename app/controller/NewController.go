package Controller

import (
	"GolangwithFrame/src/app/service"
	"GolangwithFrame/src/infrastructure/cache"
)

type Controller struct {
	service service.Service
	cache   cache.ProductCache
}

func New(service service.Service, cache cache.ProductCache) Controller {
	return Controller{
		service: service,
		cache:   cache,
	}
}
