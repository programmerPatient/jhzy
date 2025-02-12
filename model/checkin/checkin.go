package checkin

import (
	"file/model/checkinTime"
	"file/model/user"
)

// Checkin gorm实现一个打卡模型，这个模型用来记录用户打卡的信息，要记录用户每一天的打卡情况
type Checkin struct {
	ID        int    `gorm:"primaryKey"`
	UserID    int    `gorm:"column:user_id"`
	CheckinAt string `gorm:"column:checkin_at"`
	Status    int    `gorm:"column:status"`
	//关联关系
	User             user.User               `gorm:"foreignkey:id;references:user_id" json:"user"`
	CheckinAtContent checkinTime.CheckinTime `gorm:"foreignkey:checkin_at;references:checkin_at" json:"checkin_time"`
}
