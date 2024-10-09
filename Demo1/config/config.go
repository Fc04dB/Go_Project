package config

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

// InitConfig 读取配置文件
func InitConfig() {
	err := godotenv.Load() // 加载 .env 文件中的环境变量
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// InitDB 初始化数据库连接
func InitDB() {
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
}
