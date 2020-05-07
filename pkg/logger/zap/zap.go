package zap

import (
	"os"

	"github.com/nydan/glean/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type log struct {
	sl *zap.SugaredLogger
}

// Config stores the config for the logger
type Config struct {
	EnableConsole     bool
	ConsoleJSONFormat bool
	ConsoleLevel      string
	EnableFile        bool
	FileJSONFormat    bool
	FileLevel         string
	FileLocation      string
}

func getZapLevel(level string) zapcore.Level {
	switch level {
	case logger.Info:
		return zapcore.InfoLevel
	case logger.Warn:
		return zapcore.WarnLevel
	case logger.Debug:
		return zapcore.DebugLevel
	case logger.Error:
		return zapcore.ErrorLevel
	case logger.Fatal:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

func getEncoder(isJSON bool) zapcore.Encoder {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	if isJSON {
		return zapcore.NewJSONEncoder(encoderCfg)
	}
	return zapcore.NewConsoleEncoder(encoderCfg)
}

func logToConsole(cfg Config) zapcore.Core {
	level := getZapLevel(cfg.ConsoleLevel)
	writer := zapcore.Lock(os.Stdout)
	return zapcore.NewCore(getEncoder(cfg.ConsoleJSONFormat), writer, level)
}

func logToFile(cfg Config) zapcore.Core {
	level := getZapLevel(cfg.FileLevel)
	writer := zapcore.AddSync(&lumberjack.Logger{
		Filename: cfg.FileLocation,
		MaxSize:  100, // maximum filesize 100MB
		Compress: true,
		MaxAge:   28, // maximum number of days to retain old logs is 28 days
	})
	return zapcore.NewCore(getEncoder(cfg.FileJSONFormat), writer, level)
}

// NewLogger creates new zap log instance
func NewLogger(cfg Config) logger.Logger {
	cores := []zapcore.Core{}

	if cfg.EnableConsole {
		cores = append(cores, logToConsole(cfg))
	}

	if cfg.EnableFile {
		cores = append(cores, logToFile(cfg))
	}

	combinedCores := zapcore.NewTee(cores...)

	logger := zap.New(combinedCores,
		zap.AddCallerSkip(2),
		zap.AddCaller()).Sugar()

	return &log{
		sl: logger,
	}
}

func (l *log) Debugf(format string, args ...interface{}) {
	l.sl.Debugf(format, args...)
}

func (l *log) Infof(format string, args ...interface{}) {
	l.sl.Infof(format, args...)
}

func (l *log) Warnf(format string, args ...interface{}) {
	l.sl.Warnf(format, args...)
}

func (l *log) Errorf(format string, args ...interface{}) {
	l.sl.Errorf(format, args...)
}

func (l *log) Infow(format string, args ...interface{}) {
	l.sl.Infow(format, args...)
}

func (l *log) Warnw(message string, args ...interface{}) {
	l.sl.Warnw(message, args...)
}

func (l *log) Errorw(message string, args ...interface{}) {
	l.sl.Errorw(message, args...)
}

func (l *log) Fatalw(message string, args ...interface{}) {
	l.sl.Fatalw(message, args...)
}

func (l *log) WithFields(fields logger.Fields) logger.Logger {
	var f = make([]interface{}, 0)
	for k, v := range fields {
		f = append(f, k)
		f = append(f, v)
	}
	newLogger := l.sl.With(f...)
	return &log{sl: newLogger}
}
