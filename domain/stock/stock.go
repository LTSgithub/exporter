package stock

import (
	"context"
	"github.com/lits01/xiaozhan/pkg/engin"
)

type Stock struct {
	engin * engin.Engin
}

func (m * Stock)List(ctx context.Context,limit int )  {

	m.engin.GetStock()

}
