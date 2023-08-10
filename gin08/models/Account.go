package models

type Account struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Balance  int    `json:"balance"`
}

// 表示配置操作数据库的表名称
func (Account) TableName() string {
	return "account"
}
