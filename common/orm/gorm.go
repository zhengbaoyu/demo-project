package orm

import (
	"demo-project/model/demo_user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewGorm(dsn string) *gorm.DB {
	mysqlConfig := mysql.Config{
		DSN: dsn, // DSN data source name

	}
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	db.AutoMigrate(&demo_user.DemoUser{})
	if err != nil {
		panic(err)
	}
	return db
}
