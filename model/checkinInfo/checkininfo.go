package checkin

type CheckinInfo struct {
	ID        int    `gorm:"primaryKey"`
	CheckinAt string `gorm:"column:checkin_at"`
	Book      string `gorm:"column:book"`
	Chapter   string `gorm:"column:chapter"`
}
