package service

import (
	"file/model/checkin"
	"file/model/checkinTime"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	"strings"
	"time"
)

var CheckinS = &CheckinService{}

type CheckinService struct {
	BaseService
}

func (s *CheckinService) Index(c *gin.Context) {
	id := c.Query("id")
	page := c.Query("page")
	//获取当前时间的前一周
	startTime := time.Now().AddDate(0, 0, -2+7*cast.ToInt(page))
	endTime := time.Now().AddDate(0, 0, 7+7*cast.ToInt(page))
	list := checkin.OrderGetInfo(c, "checkin_at asc", [][]interface{}{
		{"user_id", "=", id},
		{"checkin_at", ">=", startTime.Format("2006-01-02")},
		{"checkin_at", " <= ", endTime.Format("2006-01-02")},
	})
	type ListItem struct {
		UserId         int
		Data           string
		CheckinContent string
		Status         int
	}
	haveAll := map[string]ListItem{}
	for _, v := range list {
		haveAll[v.CheckinAt] = ListItem{
			Data:           v.CheckinAt,
			CheckinContent: v.CheckinAtContent.Content,
			Status:         v.Status,
		}
	}
	timeList := checkinTime.OrderGetInfo(c, "checkin_at asc", [][]interface{}{
		{"checkin_at", ">=", startTime.Format("2006-01-02")},
		{"checkin_at", " <= ", endTime.Format("2006-01-02")},
	})
	all := map[string]ListItem{}
	for _, vv := range timeList {
		status := 0
		if v, ok := haveAll[vv.CheckinAt]; ok {
			status = v.Status
		}
		// 计算字符宽度，中文字符算 2，英文字符算 1
		width := 0
		for _, r := range vv.Content {
			if r > 127 { // 判断是否是非 ASCII 字符（如中文）
				width += 2
			} else {
				width++
			}
		}
		// 如果宽度不足，补全空格
		if width < 23 {
			vv.Content += strings.Repeat(" ", 23-width)
		}
		all[vv.CheckinAt] = ListItem{
			UserId:         cast.ToInt(id),
			Data:           vv.CheckinAt,
			CheckinContent: vv.Content,
			Status:         status,
		}
	}
	s.HTML(c, "checkin.tmpl", gin.H{"list": all, "id": id, "page": page})
}
func (s *CheckinService) Click(c *gin.Context) {
	timeStr := c.Query("time")
	userId := cast.ToInt(c.Query("userId"))
	if checkin.CheckinIsExist(c, userId, timeStr) {
		c.JSON(http.StatusOK, "已签到")
		return
	}
	_ = checkin.CreateByStruct(c, &checkin.Checkin{
		CheckinAt: timeStr,
		UserID:    userId,
		Status:    1,
	})
	s.Success(c, nil)
}
