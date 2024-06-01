package entity

type User struct {
	Id   int    `gorm:"id, primaryKey"`
	Name string `gorm:"name, not null"`
	Age  int    `gorm:"age, not null, default:0"`
}

func (u *User) TableName() string {
	return "users"
}
