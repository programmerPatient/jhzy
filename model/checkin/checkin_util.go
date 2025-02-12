/*
 * @Description:
 * @Author: mali
 * @Date: 2022-11-03 10:51:16
 * @LastEditTime: 2023-01-04 09:22:48
 * @LastEditors: VSCode
 * @Reference:
 */
package checkin

import (
	"context"
	"file/core/database"
	"github.com/spf13/cast"
)

/**
 * @Author: mali
 * @Func:
 * @Description: 获取单条数据
 * @Param:
 * @Return:
 * @Example:
 * @param {string} order 排序条件
 * @param {map[string]interface{}} where 条件
 * @param {...string} field
 */
func OrderGetInfo(ctx context.Context, order string, where [][]interface{}, field ...string) []Checkin {
	var app_product []Checkin
	query := database.DB.WithContext(ctx).Preload("CheckinAtContent")
	for _, v := range where {
		query.Where(cast.ToString(v[0])+" "+cast.ToString(v[1])+" ?", v[2])
	}
	query.Select(field).Order(order).Find(&app_product)
	return app_product
}

/**
 * @Author: mali
 * @Func:
 * @Description: 获取单条数据
 * @Param:
 * @Return:
 * @Example:
 * @param {string} order 排序条件
 * @param {map[string]interface{}} where 条件
 * @param {...string} field
 */
func GetInfo(ctx context.Context, where map[string]interface{}, field ...string) []Checkin {
	var app_product []Checkin
	database.DB.WithContext(ctx).Preload("User").Where(where).Select(field).Find(&app_product)
	return app_product
}

func CheckinIsExist(ctx context.Context, userId int, time string) bool {
	var count int64
	database.DB.WithContext(ctx).Where("user_id = ? and checkin_at = ?", userId, time).Count(&count)
	return count > 0
}

/**
 * @Author: mali
 * @Func:
 * @Description: 新增单条数据
 * @Param:
 * @Return:
 * @Example:
 * @param {map[string]interface{}} data
 */
func CreateByStruct(ctx context.Context, data *Checkin) error {
	db := database.DB
	result := db.WithContext(ctx).Create(data)
	if err := result.Error; err != nil {
		return err
	} else {
		return nil
	}
}
