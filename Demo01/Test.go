package models

type User struct { // 默认表名是 `users`
	Id       int
	Username string
	Age      int
	Email    string
	AddTime  int
}

func (User) TableName() string {
	return "user"
}
