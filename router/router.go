package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/yongchengchen/gf-customer-api/app/api"
	"github.com/yongchengchen/gf-customer-api/app/service"
)

func init() {
	s := g.Server()

	s.Group("/c", func(group *ghttp.RouterGroup) {
		group.Middleware(
			service.Middleware.Ctx,
			service.Middleware.CORS,
			service.Middleware.InnerAuth,
		)

		group.POST("/set", api.JcApi.SetValue)
		group.GET("/get/:key", api.JcApi.GetValue)

		group.POST("/hset", api.JcApi.HSetValue)
		group.GET("/hget/:key/:field", api.JcApi.HGetValue)

		group.POST("/push/:direction", api.JcApi.Push)
		group.POST("/maxlpush/:limit", api.JcApi.LimitLPush)

		group.POST("/push/:direction", api.JcApi.Push)
		group.POST("/maxlpush/:limit", api.JcApi.LimitLPush)

		group.GET("/lrange/:key/:from/:to", api.JcApi.LRange)
		group.GET("/llen/:key", api.JcApi.LLen)
		group.GET("/hkeys/:key", api.JcApi.HKeys)

		group.POST("/expire/:key/:ttl", api.JcApi.Expire)
		group.GET("/ttl/:key", api.JcApi.Ttl)
		group.DELETE("/del/:key", api.JcApi.Del)
		group.GET("/info", api.JcApi.Info)
	})

	// 分组路由注册方式
	// s.Group("/", func(group *ghttp.RouterGroup) {
	// 	group.Middleware(
	// 		service.Middleware.Ctx,
	// 		service.Middleware.CORS,
	// 	)
	// 	// group.ALL("/chat", api.Chat)
	// 	group.ALL("/user", api.User)
	// 	group.Group("/", func(group *ghttp.RouterGroup) {
	// 		group.Middleware(service.Middleware.Auth)
	// 		group.ALL("/user/profile", api.User.Profile)
	// 	})
	// })
}
