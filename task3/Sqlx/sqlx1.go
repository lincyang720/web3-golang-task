// 假设你已经使用Sqlx连接到一个数据库，
// 并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
// 要求 ：
// 编写Go代码，使用Sqlx查询 employees
// 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
// 编写Go代码，使用Sqlx查询 employees
// 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

type Employee struct {
	Id         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

type Book struct {
	Id     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func main() {
	db, err := sqlx.Connect("mysql", "root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接数据库失败：", err)
		return
	}
	defer db.Close()

	// 查询所有部门位技术部员工信息
	// var employees []Employee
	// db.Select(&employees, "SELECT * FROM employees WHERE department = ?", "技术部")

	// var employee Employee
	// db.Get(&employee, "SELECT * FROM employees ORDER BY salary DESC LIMIT 1")

	// 确保 books 表存在
	db.MustExec("CREATE TABLE IF NOT EXISTS books (id INT AUTO_INCREMENT PRIMARY KEY, title VARCHAR(255), author VARCHAR(255), price DECIMAL(10,2))")

	var books []Book
	err = db.Select(&books, "SELECT * FROM books WHERE price > ?", 50)
	if err != nil {
		fmt.Println("查询失败：", err)
		return
	}

}
