package user

// User gorm 用户表
type User struct {
	//自增id
	ID       int    `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(255);not null"`
	Password string `gorm:"type:varchar(255);not null"`
}
