package logger

import (
	"fake-server/config"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var Logger *FMLogger

const (
	DefaultLogRotateSize  = 100
	DefaultLogBackupCount = 100
	DefaultLogMaxAge      = 100
)

type FMLogger struct {
	logger *zap.SugaredLogger
}

// Debugf print debug message
func (l *FMLogger) Debugf(format string, a ...interface{}) {
	l.logger.Debugf(format, a...)
}

// Infof print info message
func (l *FMLogger) Infof(format string, a ...interface{}) {
	l.logger.Infof(format, a...)
}

// Warnf print warning message
func (l *FMLogger) Warnf(format string, a ...interface{}) {
	l.logger.Warnf(format, a...)
}

// Errorf print error message
func (l *FMLogger) Errorf(format string, a ...interface{}) {
	l.logger.Errorf(format, a...)
}

// WithField tag logger with key and value
func (l *FMLogger) WithField(k string, v interface{}) *FMLogger {
	logger := &FMLogger{
		logger: l.logger.With(k, v),
	}

	return logger
}

func Init(fields ...zap.Field) (*FMLogger, error) {
	// 创建日志配置
	var encoderConfig = zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// 创建Console输出Core
	consoleEncoder := zapcore.NewJSONEncoder(encoderConfig)
	atomicLevel := zap.NewAtomicLevel()

	consoleCore := zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), atomicLevel)
	pid := os.Getpid()
	filePath := config.GlobalConfig.LogPath + fmt.Sprintf("/run_%d.log", pid)
	// 创建日志输出Core
	fmt.Printf("Init logger setting: MaxSize: %d, MaxAge: %d, MaxBackups: %d, logPath: %s\n",
		DefaultLogRotateSize, DefaultLogMaxAge, DefaultLogBackupCount, filePath)
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    DefaultLogRotateSize,  // megabytes
		MaxAge:     DefaultLogBackupCount, // backup log files 7 Days
		MaxBackups: DefaultLogMaxAge,      // backup log files 100
		Compress:   true,
	})
	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		fileWriter,
		atomicLevel,
	)
	cores := zapcore.NewTee(fileCore, consoleCore)
	log := zap.New(cores)
	log = log.WithOptions(
		zap.ErrorOutput(os.Stdout), // error message output to stdout
		zap.AddCaller(),            // add function caller info to log
		zap.AddCallerSkip(1),       // make stack having right depth to get function call
		zap.Fields(fields...),      // add common log info, like local_ip
	)
	logger := FMLogger{
		logger: log.Sugar(),
	}
	return &logger, nil
}
