package log

import (
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var LogZap *zap.SugaredLogger

func InitLog(level string, filePath string) {
	hook := lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     3,
		Compress:   true,
	}
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "file",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	var writers []zapcore.WriteSyncer
	writers = append(writers, os.Stdout)
	writers = append(writers, zapcore.AddSync(&hook))
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(getLevel(level))
	core := zapcore.NewCore(
		//zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(writers...),
		atomicLevel,
	)
	caller := zap.AddCaller()
	development := zap.Development()
	LogZap = zap.New(core, caller, development, zap.AddCallerSkip(1)).Sugar()
}

func getLevel(level string) (l zapcore.Level) {
	switch level {
	case "debug":
		l = zap.DebugLevel
	case "info":
		l = zap.InfoLevel
	case "warn":
		l = zap.WarnLevel
	case "error":
		l = zap.ErrorLevel
	case "panic":
		l = zap.PanicLevel
	case "fatal":
		l = zap.FatalLevel
	default:
		l = zap.InfoLevel
	}
	return
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		Info(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func Debug(args ...interface{}) {
	LogZap.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	LogZap.Debugf(template, args...)
}

func Info(args ...interface{}) {
	LogZap.Info(args...)
}

func Infof(template string, args ...interface{}) {
	LogZap.Infof(template, args...)
}

func Warn(args ...interface{}) {
	LogZap.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	LogZap.Warnf(template, args...)
}

func Error(args ...interface{}) {
	LogZap.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	LogZap.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	LogZap.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	LogZap.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	LogZap.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	LogZap.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	LogZap.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	LogZap.Fatalf(template, args...)
}
