// Code generated by sqlc. DO NOT EDIT.

package generated

import (
	"time"
)

type AlertConfig struct {
	UserID     string  `db:"user_id" json:"userID"`
	StockCode  string  `db:"stock_code" json:"stockCode"`
	NotifyType string  `db:"notify_type" json:"notifyType"`
	Price      float64 `db:"price" json:"price"`
	ID         string  `db:"id" json:"id"`
	Deadline   int64   `db:"deadline" json:"deadline"`
	CreateTime int64   `db:"create_time" json:"createTime"`
	Desc       string  `db:"desc" json:"desc"`
}

// app状态表
type App struct {
	AppStatus string `db:"app_status" json:"appStatus"`
}

type Stock struct {
	Code       string `db:"code" json:"code"`
	Name       string `db:"name" json:"name"`
	Type       string `db:"type" json:"type"`
	CreateTime string `db:"create_time" json:"createTime"`
	UpdateTime string `db:"update_time" json:"updateTime"`
}

// 用户表
type User struct {
	// 用户id
	ID int32 `db:"id" json:"id"`
	// 用户名
	Username string `db:"username" json:"username"`
	// 密码
	Password string `db:"password" json:"password"`
	// 描述信息
	Description string `db:"description" json:"description"`
	// 创建时间
	CreateTime time.Time `db:"create_time" json:"createTime"`
}
