package Struct

type User struct {
	ID    string `json:"id"`    // 主键ID
	Name  string `json:"name"`  // 姓名
	Class string `json:"class"` // 班级
	Age   string `json:"age"`   // 年龄
}

type Msg struct {
	ID     string `json:"id"`     // 客户端ID
	Data   string `json:"data"`   // 数据
	Status bool   `json:"status"` // 状态
	Next   string `json:"next"`   // 下一个组件名
}

func (User) TableName() string {
	return "user"
}
