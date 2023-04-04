package svc

import (
	"demo-project/app/internal/config"
	"demo-project/app/internal/middleware"
	"demo-project/common/orm"
	"demo-project/model/demo_user"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config        config.Config
	Orm           *gorm.DB
	Auth          rest.Middleware
	DemoUserModel demo_user.DemoUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	mysqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:        c,
		Orm:           orm.NewGorm(c.Mysql.DataSource),
		Auth:          middleware.NewAuthMiddleware().Handle,
		DemoUserModel: demo_user.NewDemoUserModel(mysqlConn),
	}
}
