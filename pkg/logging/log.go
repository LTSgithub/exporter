package logging

import (
	"fmt"
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"github.com/mattn/go-isatty"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() *logr.Logger {
	var (
		logger *zap.Logger
		err    error
		config zap.Config
		l      zapcore.Level
	)

	if err = l.UnmarshalText([]byte("info")); err != nil {
		//nolint
		l.Set("INFO")
	}
	if isatty.IsTerminal(os.Stdout.Fd()) {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()
	}
	// 设置时间格式
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 级别序列化为大写
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	config.Level.SetLevel(l)

	//config
	logger, err = config.Build()
	if err != nil {
		panic(fmt.Sprintf("Failed to build logger: %v", err))
	}

	log := zapr.NewLogger(logger).WithName(string("main"))
	return &log
}
