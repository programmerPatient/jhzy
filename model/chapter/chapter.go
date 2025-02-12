package chapter

import "file/model/Book"

type Chapter struct {
	ID       int       `gorm:"primaryKey"`
	BookId   int       `gorm:"column:book_id"` // 书卷id
	Chapter  int       `gorm:"column:chapter"` // 章
	Section  int       `gorm:"column:section"` // 章节
	Content  string    `gorm:"column:content"` // 章节内容
	BookInfo Book.Book `gorm:"foreignkey:id;references:book_id" json:"book"`
}
