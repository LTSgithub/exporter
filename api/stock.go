package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lits01/xiaozhan/app"
	typing "github.com/lits01/xiaozhan/type"
)

type Stock struct {
	Base
	stock * app.Stock
}

func (m *Stock)GetStock(ctx *gin.Context) {
	req := &typing.GetStockRequest{}
	if err :=m.UnmarshalRequest(req);err != nil {
		m.Response(nil,err)
	}
	m.Response(m.stock.Engin.GetStock(ctx.Request.Context(),req.Code))
}

func (m * Stock)GetStocks(ctx * gin.Context)  {

}