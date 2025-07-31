package object

/**
 * 使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
 * 为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
 */
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Employee struct {
	Person
	EmployeeID int `json:"employee_id"`
}

func (e Employee) PrintInfo() string {
	return "Employee Info: Name: " + e.Name + ", Age: " + string(e.Age) + ", EmployeeID: " + string(e.EmployeeID)
}
