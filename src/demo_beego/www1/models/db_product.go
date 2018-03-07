package models

type Product struct {
	Id   int
	name  string
	count  int16
	ctime   string
	User *User `orm:"reverse(one)"` // 设置反向关系(可选)
}
