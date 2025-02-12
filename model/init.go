package model

import (
	"file/core/database"
	"file/model/checkin"
	"file/model/checkinTime"
	"file/model/user"
)

func Init() {
	_ = database.DB.Table("checkin_times").AutoMigrate(&checkinTime.CheckinTime{})
	_ = database.DB.Table("checkins").AutoMigrate(&checkin.Checkin{})
	_ = database.DB.Table("users").AutoMigrate(&user.User{})
}
