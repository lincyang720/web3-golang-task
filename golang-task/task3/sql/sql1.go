package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 基本CRUD操作
type Student struct {
	Id    int    `gorm:"primary_key;AUTO_INCREMENT"`
	Name  string `gorm:"type:varchar(100);not null"`
	Age   int    `gorm:"type:int;not null"`
	Grade string `gorm:"type:varchar(50);not null"`
}

//要求 ：
// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。

// 事务语句
type Accounts struct {
	Id     int     `gorm:"primary_key;AUTO_INCREMENT"`
	Blance float64 `gorm:"type:decimal(10,2);not null"`
}

type Transaction struct {
	Id            int     `gorm:"primary_key;AUTO_INCREMENT"`
	FromAccountId int     `gorm:"type:int;not null"`
	ToAccountId   int     `gorm:"type:int;not null"`
	Amount        float64 `gorm:"type:decimal(10,2);not null"`
}

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// db.AutoMigrate(&Student{})
	// // 插入新记录
	// db.Create(&Student{Name: "张三", Age: 20, Grade: "三年级"})
	// // 查询所有年龄大于18岁的学生信息
	// var results []Student
	// db.Where("age > ?", 18).Find(&results)
	// // 更新姓名为 "张三" 的学生年级为 "四年级"
	// db.Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级")
	// // 删除年龄小于15岁的学生记录
	// db.Where("age < ?", 15).Delete(&Student{})
	db.AutoMigrate(&Accounts{})
	db.AutoMigrate(&Transaction{})

	//编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。
	// 在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，
	// 向账户 B 增加 100 元，
	// 并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
	db.Transaction(func(tx *gorm.DB) error {
		var accountA Accounts
		var accountB Accounts
		if err := tx.First(&accountA, 1).Error; err != nil {
			return err
		}
		if err := tx.First(&accountB, 2).Error; err != nil {
			return err
		}
		if accountA.Blance < 100 {
			return gorm.ErrRecordNotFound // 余额不足，回滚事务
		}
		accountA.Blance -= 100
		accountB.Blance += 100
		if err := tx.Save(&accountA).Error; err != nil {
			return err
		}
		if err := tx.Save(&accountB).Error; err != nil {
			return err
		}

		if err := tx.Create(&Transaction{FromAccountId: accountA.Id, ToAccountId: accountB.Id, Amount: 100}).Error; err != nil {
			return err
		}
		// 提交事务
		return tx.Commit().Error
	})
}
