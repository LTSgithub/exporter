package api

import (
	"github.com/gin-gonic/gin"
	typing "github.com/lits01/xiaozhan/type"
	"path/filepath"
)

func InitStockApi(g * gin.Engine)  {
	g.POST(filepath.Join(version,"/stock/get"),stock.GetStock)
	g.POST(filepath.Join(version,"/stocks/get"),stock.GetStocks)
}

type Stock struct {
	Base
}

func (m *Stock)GetStock(ctx *gin.Context) {
	req := &typing.GetStockRequest{}
	if err :=m.UnmarshalRequest(req);err != nil {
		m.Response(nil,err)
	}
	m.Response(m.Engin.GetStock(ctx.Request.Context(),req.Code))
}

func (m * Stock)GetStocks(ctx * gin.Context)  {

}