package initialize

import (
	"github.com/hespecial/gin-mall/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
	"time"
)

var (
	level   zapcore.Level // zap 日志等级
	options []zap.Option  // zap 配置项
)

func InitLogger() *zap.Logger {
	// 设置日志等级
	setLogLevel()
	options = append(options, zap.WithCaller(global.Config.Log.ShowLine))
	// 初始化 zap
	return zap.New(getZapCore(), options...)
}

func setLogLevel() {
	switch global.Config.Log.Level {
	case "debug":
		level = zap.DebugLevel
		options = append(options, zap.AddStacktrace(level))
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
		options = append(options, zap.AddStacktrace(level))
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
}

// 扩展 Zap
func getZapCore() zapcore.Core {
	var encoder zapcore.Encoder
	// 调整编码器默认配置
	encoderConfig := zap.NewProductionEncoderConfig()
	// 调用行显示全路径
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("[" + "2006-01-02 15:04:05" + "]"))
	}
	encoderConfig.EncodeLevel = func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(strings.ToUpper(l.String()))
	}
	// 设置编码器
	switch global.Config.Log.Format {
	case "json":
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	case "console":
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	default:
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}
	return zapcore.NewCore(encoder, getLogWriter(), level)
}

// 使用 lumberjack 作为日志写入器
func getLogWriter() zapcore.WriteSyncer {
	switch global.Config.Server.Level {
	case "debug":
		return zapcore.AddSync(os.Stdout)
	case "release":
		f := &lumberjack.Logger{
			Filename:   global.Config.Log.Dir + "/" + global.Config.Log.Filename,
			MaxSize:    global.Config.Log.MaxSize,
			MaxBackups: global.Config.Log.MaxBackups,
			MaxAge:     global.Config.Log.MaxAge,
			Compress:   global.Config.Log.Compress,
		}
		return zapcore.AddSync(f)
	default:
		return zapcore.AddSync(os.Stdout)
	}
}
