package routes

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/context"
	"github.com/zidni722/pawoon-product/app/web/controllers"
	"github.com/zidni722/pawoon-product/bootstrap"
	"github.com/zidni722/pawoon-product/config"
)

type Route struct {
	Config      *config.Configuration
	CorsHandler context.Handler
}

func NewRoute(config *config.Configuration) *Route {
	return &Route{
		Config: config,
		CorsHandler: cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowCredentials: true,
			AllowedHeaders:   []string{"*"},
		}),
	}
}

func (r *Route) Configure(b *bootstrap.Bootstrapper) {
	b.Get("/", controllers.GetHomeHandler)

	// repositories
	// itemRequestRepository := impl.NewItemRepositoryImpl()
	// orderRequestRepository := impl.NewOrderRepositoryImpl()
	// orderDetailRequestRepository := impl.NewOrderDetailRepositoryImpl()

	// services
	// itemService := impl2.NewItemServiceImpl(itemRequestRepository)
	// orderService := impl2.NewOrderServiceImpl(orderRequestRepository)
	// orderDetailService := impl2.NewOrderDetailServiceImpl(orderDetailRequestRepository)

	// v1 := b.Party("/v1", r.CorsHandler).AllowMethods(iris.MethodOptions)
	// {
	// items := v1.Party("/items")
	// {
	// 	itemController := controllers.NewItemController(r.Config.Database.DB, itemService)
	// 	items.Get("/", itemController.GetIndexHandler)
	// }
	// }
}
