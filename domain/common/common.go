package common

import (
	"context"
	"github.com/lits01/xiaozhan/config"
	"time"
)

func NewDefaultContext()(context.Context, context.CancelFunc)   {
	return context.WithTimeout(context.Background(), config.RequestTimeout*time.Second)
}
