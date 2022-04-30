package database

import (
	"fmt"

	"bamboo-api/app/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var bambooDBClient *gorm.DB

func InitMysql() {
	conf := config.GetConfig()
	dns := fmt.Sprintf("%v:%v@tcp(%v:%v)/bamboo?charset=utf8mb4&parseTime=True&loc=Local",
		conf.GetString("database.mysql.username"),
		conf.GetString("database.mysql.password"),
		conf.GetString("database.mysql.url"),
		conf.GetString("database.mysql.port"))
	client, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	bambooDBClient = client

	// Migrate the schema
	//db.AutoMigrate(&Product{})

	// Create
	//db.Create(&Product{Code: "D42", Price: 100})
	//
	//// Read
	//var product Product
	//db.First(&product, 1) // find product with integer primary key
	//db.First(&product, "code = ?", "D42") // find product with code D42
	//
	//// Update - update product's price to 200
	//db.Model(&product).Update("Price", 200)
	//// Update - update multiple fields
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//
	//// Delete - delete product
	//db.Delete(&product, 1)
}
func GetDBClient() *gorm.DB {
	return bambooDBClient
}
