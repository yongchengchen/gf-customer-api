package boot

import (
	_ "github.com/yongchengchen/gf-customer-api/packed"
	_ "github.com/yongchengchen/gf-customer-api/app/model"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/swagger"
)

// 用于应用初始化。
func init() {
	s := g.Server()
	var customers []model.Customer
	err := g.DB().Model("customers").Scan(&customers)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(customers)
	s.Plugin(&swagger.Swagger{})
}
