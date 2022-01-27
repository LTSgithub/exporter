package engin

import (
	"context"
	"time"
)

func NewDefaultContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), config.RequestTimeout*time.Second)
}
