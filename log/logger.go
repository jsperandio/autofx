package log

import (
	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var instance *zap.SugaredLogger

func Init(logLevel *string) {
	var (
		ll  zapcore.Level
		err error
	)

	logCfg := zap.NewDevelopmentEncoderConfig()
	logCfg.EncodeTime = nil
	logCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	if logLevel != nil {
		ll, err = zapcore.ParseLevel(*logLevel)
		if err != nil {
			ll = zapcore.ErrorLevel
		}
	}

	instance = zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(logCfg),
		zapcore.AddSync(colorable.NewColorableStdout()),
		ll,
	)).Sugar()
}

func GetLogger() *zap.SugaredLogger {
	return instance
}
