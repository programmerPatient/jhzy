/*
 * @Description:
 * @Author: mali
 * @Date: 2022-11-03 10:51:16
 * @LastEditTime: 2023-01-04 09:22:48
 * @LastEditors: VSCode
 * @Reference:
 */
package checkinTime

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
func OrderGetInfo(ctx context.Context, order string, where [][]interface{}, field ...string) []CheckinTime {
	var app_product []CheckinTime
	query := database.DB.WithContext(ctx)
	for _, v := range where {
		query = query.Where(cast.ToString(v[0])+" "+cast.ToString(v[1])+" ?", v[2])
	}
	query = query.Select(field).Order(order)
	// 打印 SQL 语句
	query = query.Debug()
	query.Find(&app_product)
	return app_product
}
