package engin

import (
	"context"
	"time"

	"github.com/lits01/xiaozhan/config"
)

func NewDefaultContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), config.RequestTimeout*time.Second)
}
