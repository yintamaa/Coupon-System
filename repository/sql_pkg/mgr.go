package sql_pkg

import (
	"coupon_system/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type sqlManager struct {
	dsn string
	db  *gorm.DB
}

var sqlMgr *sqlManager

func InitSqlManager() error {
	// dsn := "root:123456@tcp(127.0.0.1:3306)/blogok?charset=utf8mb4&parseTime=True&loc=Local"
	// 建立数据库连接
	conf := config.GetCouponConfig()
	db, err := gorm.Open(mysql.Open(conf.MysqlDSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移表结构
	db.AutoMigrate(&Coupon{})

	// // 增加数据
	// db.Create(&Product{Code: "D42", Price: 100})

	// // 查找数据
	// var product Product
	// db.First(&product, 1)                 // find product with integer primary key
	// db.First(&product, "code = ?", "D42") // find product with code D42

	// // 更新数据 - update product's price to 200
	// db.Model(&product).Update("Price", 200)
	// // 更新数据 - update multiple fields
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // 删除数据 - delete product
	// db.Delete(&product, 1)

	sqlMgr = &sqlManager{
		db:  db,
		dsn: conf.MysqlDSN,
	}

	return err
}

func GetSqlManager() *sqlManager {
	return sqlMgr
}
