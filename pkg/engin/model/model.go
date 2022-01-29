package model

import typing "github.com/lits01/xiaozhan/type"

type Stock struct {
	Code  string
	Name  string
	Type  string
	Price float32
}

type Status struct {
	UpdateTime int64
	Price      float32
	Prices     []*typing.TV
}

type StockDetail struct {
	Stock
	RealTime *Status
	Minutes  *Status
	Days     *Status
	Weeks    *Status
	Months   *Status
}
