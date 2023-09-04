package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

func Init() *zap.SugaredLogger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Format(time.DateTime))
	}

	infoLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level <= zap.InfoLevel
	})
	errLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level > zap.InfoLevel
	})

	//infoFile, _ := os.OpenFile("/info.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	//errFile, _ := os.OpenFile("/err.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	logger := zap.New(
		zapcore.NewTee(
			zapcore.NewCore(
				zapcore.NewJSONEncoder(config.EncoderConfig),
				zapcore.AddSync(os.Stdout),
				infoLevel,
			),
			zapcore.NewCore(
				zapcore.NewJSONEncoder(config.EncoderConfig),
				zapcore.AddSync(os.Stdout),
				errLevel,
			),
		),
		zap.AddCaller(),
		zap.AddStacktrace(errLevel),
	)
	return logger.Sugar()
}
