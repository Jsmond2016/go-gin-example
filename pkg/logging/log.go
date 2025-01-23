package logging

import (
	"fmt"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/EDDYCJY/go-gin-example/pkg/file"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
)

var logger *zap.Logger

// Setup initialize the log instance
func Setup() {
	logPath := getLogFilePath()
	fileName := getLogFileName()
	filePath := filepath.Join(logPath, fileName)

	// 确保日志目录存在
	if err := file.IsNotExistMkDir(logPath); err != nil {
		panic(fmt.Sprintf("create log directory '%s' failed: %v", logPath, err))
	}

	// 日志轮转配置
	hook := &lumberjack.Logger{
		Filename:   filePath, // 日志文件路径
		MaxSize:    500,      // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 3,        // 日志文件最多保存多少个备份
		MaxAge:     28,       // 文件最多保存多少天
		Compress:   true,     // 是否压缩
	}
	writeSyncer := zapcore.AddSync(hook)

	// 编码器配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     timeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	if setting.ServerSetting.RunMode == "debug" {
		atomicLevel.SetLevel(zap.DebugLevel)
	} else {
		atomicLevel.SetLevel(zap.InfoLevel)
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // 编码器配置
		writeSyncer,                           // 打印到文件
		atomicLevel,                           // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	filed := zap.Fields(zap.String("serviceName", "go-gin-example"))
	// 构造日志
	logger = zap.New(core, caller, development, filed)
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// Debug output logs at debug level
func Debug(v ...interface{}) {
	logger.Sugar().Debug(v...)
}

// Info output logs at info level
func Info(v ...interface{}) {
	logger.Sugar().Info(v...)
}

// Warn output logs at warn level
func Warn(v ...interface{}) {
	logger.Sugar().Warn(v...)
}

// Error output logs at error level
func Error(v ...interface{}) {
	logger.Sugar().Error(v...)
}

// Fatal output logs at fatal level
func Fatal(v ...interface{}) {
	logger.Sugar().Fatal(v...)
}

// Close flushes any buffered log entries
func Close() error {
	return logger.Sync()
}
