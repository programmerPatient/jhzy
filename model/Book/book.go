package Book

type Book struct {
	Id       int    `gorm:"primaryKey"`
	BookName string `gorm:"column:book_name"`
}
