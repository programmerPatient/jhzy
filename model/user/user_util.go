package user

import (
	"context"
	"file/core/database"
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
func OrderGetInfo(ctx context.Context, order string, where map[string]interface{}, field ...string) []User {
	var app_product []User
	database.DB.WithContext(ctx).Preload("User").Where(where).Select(field).Order(order).Find(&app_product)
	return app_product
}
