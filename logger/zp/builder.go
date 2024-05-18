package zp

import (
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"

	"github.com/cc-cheunggg/ioc-golang/logger/common"
)

var config common.Config

var Factory ZapFactory = func(c common.Config) {
	config = c
}

type ZapFactory func(config common.Config)

func (f ZapFactory) cg() common.Config {
	return config
}

func (f ZapFactory) zapEncodeLevel() zapcore.LevelEncoder {
	switch {
	case f.cg().Encoder == "LowercaseLevelEncoder": // 小写编码器(默认)
		return zapcore.LowercaseLevelEncoder
	case f.cg().Encoder == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		return zapcore.LowercaseColorLevelEncoder
	case f.cg().Encoder == "CapitalLevelEncoder": // 大写编码器
		return zapcore.CapitalLevelEncoder
	case f.cg().Encoder == "CapitalColorLevelEncoder": // 大写编码器带颜色
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}

func (f ZapFactory) transportLevel() zapcore.Level {
	switch f.cg().Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.WarnLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}

func (f ZapFactory) customTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(f.cg().Prefix + t.Format("2006-01-02 - 15:04:05.000"))
}

func (f ZapFactory) writeSyncer(level string) (zapcore.WriteSyncer, error) {

	dateStr := time.Now().Format("2006-01-02")
	if ok, _ := pathExists(dateStr); !ok { // 判断是否有日期文件夹
		_ = os.Mkdir(path.Join(f.cg().Director, dateStr), os.ModePerm)
	}
	fileWriter, err := rotatelogs.New(
		path.Join(f.cg().Director, dateStr, level+".log"),
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithMaxAge(time.Duration(f.cg().MaxAge)*24*time.Hour), // 日志留存时间
		rotatelogs.WithRotationTime(time.Hour*24),
	)

	if f.cg().Std {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}

	return zapcore.AddSync(fileWriter), err
}

func (f ZapFactory) levelPriority(level zapcore.Level) zap.LevelEnablerFunc {
	switch level {
	case zapcore.DebugLevel:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	case zapcore.InfoLevel:
		return func(level zapcore.Level) bool { // 日志级别
			return level == zap.InfoLevel
		}
	case zapcore.WarnLevel:
		return func(level zapcore.Level) bool { // 警告级别
			return level == zap.WarnLevel
		}
	case zapcore.ErrorLevel:
		return func(level zapcore.Level) bool { // 错误级别
			return level == zap.ErrorLevel
		}
	case zapcore.DPanicLevel:
		return func(level zapcore.Level) bool { // dpanic级别
			return level == zap.DPanicLevel
		}
	case zapcore.PanicLevel:
		return func(level zapcore.Level) bool { // panic级别
			return level == zap.PanicLevel
		}
	case zapcore.FatalLevel:
		return func(level zapcore.Level) bool { // 终止级别
			return level == zap.FatalLevel
		}
	default:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	}
}

func (f ZapFactory) encoder(l zapcore.Level, level zap.LevelEnablerFunc) zapcore.Core {
	config := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  f.cg().StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    f.zapEncodeLevel(),
		EncodeTime:     f.customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	var encoder zapcore.Encoder
	if f.cg().Format == "json" {
		encoder = zapcore.NewJSONEncoder(config)
	} else {
		encoder = zapcore.NewConsoleEncoder(config)
	}
	if writer, err := f.writeSyncer(l.String()); err != nil { // 使用file-rotatelogs进行日志分割
		panic(err)
	} else {
		return zapcore.NewCore(encoder, writer, level)
	}

}

func (f ZapFactory) zapCores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)
	for level := f.transportLevel(); level <= zapcore.FatalLevel; level++ {
		cores = append(cores, f.encoder(level, f.levelPriority(level)))
	}
	return cores
}

func pathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("duplicate file exists")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (f ZapFactory) NewZap() *zap.Logger {
	if ok, _ := pathExists(f.cg().Director); !ok { // 判断是否有Director文件夹
		_ = os.Mkdir(f.cg().Director, os.ModePerm)
	}
	cores := f.zapCores()
	logger := zap.New(zapcore.NewTee(cores...))

	if f.cg().ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
