package api

import (
	"github.com/gin-gonic/gin"
	"path/filepath"
)

var (
	stock *Stock
	version  = "/v1"
)

func AddStock(g * gin.Engine)  {
	g.POST(filepath.Join(version,"/stock/create"),stock.GetStock)
}