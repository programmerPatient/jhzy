package checkinTime

// CheckinTime gorm实现一个打卡模型，这个模型用来记录用户打卡的信息，要记录用户每一天的打卡情况
type CheckinTime struct {
	CheckinAt string `gorm:"primaryKey;column:checkin_at"` //年份
	Content   string `gorm:"column:content"`               //章节内容
}
