package dao

import (
	"fmt"
	"go_code/ginStudy/gin01_b/bubble/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	dsn := "root:2004213@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	// 获取底层 *sql.DB 对象
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}

	// 执行 Ping 操作
	err = sqlDB.Ping()
	if err != nil {
		return err
	}

	return nil
}

func CloseMySQL() {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			fmt.Printf("获取底层数据库实例失败: %v\n", err)
			return
		}
		if err := sqlDB.Close(); err != nil {
			fmt.Printf("关闭数据库连接失败: %v\n", err)
		} else {
			fmt.Println("数据库连接已关闭")
		}
	}
}

// InitModel 初始化数据库表结构
func InitModel() error {
	// 自动迁移 Todo 表
	if err := DB.AutoMigrate(&models.Todo{}); err != nil {
		return err
	}
	return nil
}
