package zlogger

import (
	"context"
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Field = zap.Field

// Logger methods interface
type Logger interface {
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	DPanic(args ...interface{})
	DPanicf(template string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
	// log with field
	InfoWithField(msg string, f ...Field)
}

var GLogger *logger = &logger{}

// Logger
type logger struct {
	sugarLogger *zap.SugaredLogger
	l           *zap.Logger
	// zapcore.Core
	Service  string
	logLevel int
	// Path    string
}

func configure() zapcore.WriteSyncer {

	return zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(os.Stdout),
		// zapcore.AddSync(w),
	)
}

// App Logger constructor
func NewLogger(mode, level, format, service string) Logger {
	logLevel, exist := loggerLevelMap[level]
	if !exist {
		logLevel = zapcore.DebugLevel
	}

	// todo
	// ap
	var encoderCfg zapcore.EncoderConfig
	if mode == "development" {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}

	encoderCfg.LevelKey = "LEVEL"
	encoderCfg.CallerKey = "CALLER"
	encoderCfg.TimeKey = "TIME"
	encoderCfg.NameKey = "NAME"
	encoderCfg.MessageKey = "MESSAGE"
	encoderCfg.EncodeDuration = zapcore.NanosDurationEncoder
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	// encoderCfg.EncodeCaller = zapcore.FullCallerEncoder
	// encoderCfg.LineEnding = zapcore.DefaultLineEnding

	var encoder zapcore.Encoder
	if format == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}
	if service != "" {
		encoder.AddString("service", service)
		// encoder.AddString("traceId", utils.GetTraceID(12))
	}
	core := zapcore.NewCore(encoder, configure(), zap.NewAtomicLevelAt(logLevel))
	loggerZap := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	sugarLogger := loggerZap.Sugar()
	tempLog := &logger{
		sugarLogger: sugarLogger,
		l:           loggerZap,
		logLevel:    int(logLevel),
		// Core:        core,
	}
	GLogger = tempLog
	return tempLog
}

// For mapping config logger to app logger levels
var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

// Logger methods

func (l *logger) Debug(args ...interface{}) {
	l.sugarLogger.Debug(args...)
}

func (l *logger) Debugf(template string, args ...interface{}) {
	l.sugarLogger.Debugf(template, args...)
}

func (l *logger) Info(args ...interface{}) {
	l.sugarLogger.Info(args...)
}

func (l *logger) Infof(template string, args ...interface{}) {
	l.sugarLogger.Infof(template, args...)
}

func (l *logger) Warn(args ...interface{}) {
	l.sugarLogger.Warn(args...)
}

func (l *logger) Warnf(template string, args ...interface{}) {
	l.sugarLogger.Warnf(template, args...)
}

func (l *logger) Error(args ...interface{}) {
	l.sugarLogger.Error(args...)
}

func (l *logger) Errorf(template string, args ...interface{}) {
	l.sugarLogger.Errorf(template, args...)
}

func (l *logger) DPanic(args ...interface{}) {
	l.sugarLogger.DPanic(args...)
}

func (l *logger) DPanicf(template string, args ...interface{}) {
	l.sugarLogger.DPanicf(template, args...)
}

func (l *logger) Panic(args ...interface{}) {
	l.sugarLogger.Panic(args...)
}

func (l *logger) Panicf(template string, args ...interface{}) {
	l.sugarLogger.Panicf(template, args...)
}

func (l *logger) Fatal(args ...interface{}) {
	l.sugarLogger.Fatal(args...)
}

func (l *logger) Fatalf(template string, args ...interface{}) {
	l.sugarLogger.Fatalf(template, args...)
}

// log with field
// use zap.Logger
func (l *logger) InfoWithField(msg string, f ...Field) {
	l.l.Info(msg, f...)
}

func GetFieldsTrace(traceId string) Field {
	return zap.String("traceId", traceId)
}

func GetFieldsReqId(requestId string) Field {
	return zap.String("requestId", requestId)
}

func GetFieldsKafkaMessageType(messageType string) Field {
	return zap.String("type", messageType)
}

func GetFieldsWorkerID(workerID string) Field {
	return zap.String("workerID", workerID)
}

func GetFieldsFunctions(funcName string) Field {
	return zap.String("functions", funcName)
}

// logging
type contextKey string

const (
	ContextKeyRequestID  contextKey = "reqId"
	ContextKeyDocumentId contextKey = "documentId"
)

var fieldsLog = []contextKey{ContextKeyRequestID, ContextKeyDocumentId}

func (l *logger) InfoCtxWithField(ctx context.Context, msg string, f ...Field) {
	newLogger := l.l
	if ctx != nil {
		if ctxRqId, ok := ctx.Value(ContextKeyRequestID).(string); ok {
			newLogger = l.l.With(zap.String(string(ContextKeyRequestID), ctxRqId))
		}
		if userId, ok := ctx.Value("user").(string); ok {
			newLogger = l.l.With(zap.String("user", userId))
		}
	}
	newLogger.Info(msg, f...)
}

func (l *logger) DebugCtxWithField(ctx context.Context, msg string, f ...Field) {
	newLogger := l.l
	if ctx != nil {
		if ctxRqId, ok := ctx.Value(ContextKeyRequestID).(string); ok {
			newLogger = l.l.With(zap.String(string(ContextKeyRequestID), ctxRqId))
		}
		if userId, ok := ctx.Value("user").(string); ok {
			newLogger = l.l.With(zap.String("user", userId))
		}
	}
	newLogger.Debug(msg, f...)
}
func (l *logger) ErrorCtxWithField(ctx context.Context, msg string, f ...Field) {
	newLogger := l.l
	if ctx != nil {
		if ctxRqId, ok := ctx.Value(ContextKeyRequestID).(string); ok {
			newLogger = l.l.With(zap.String(string(ContextKeyRequestID), ctxRqId))
		}
		if userId, ok := ctx.Value("user").(string); ok {
			newLogger = l.l.With(zap.String("user", userId))
		}
	}
	newLogger.Error(msg, f...)
}

func (l *logger) logWithContext(ctx context.Context) *zap.Logger {
	newLogger := l.l

	for _, field := range fieldsLog {
		if value := ctx.Value(field); value != nil {
			newLogger = newLogger.With(zap.Any(string(field), value))
		}
	}

	return newLogger
}

func (l *logger) InfoWithContext(ctx context.Context, msg string, args ...interface{}) {
	newLogger := l.logWithContext(ctx)
	newLogger.Info(fmt.Sprintf(msg, args...))
}

func (l *logger) ErrorWithContext(ctx context.Context, msg string, args ...interface{}) {
	newLogger := l.logWithContext(ctx)
	newLogger.Error(fmt.Sprintf(msg, args...))
}

func (l *logger) DebugWithContext(ctx context.Context, msg string, args ...interface{}) {
	newLogger := l.logWithContext(ctx)
	newLogger.Debug(fmt.Sprintf(msg, args...))
}

func (l *logger) WarnWithContext(ctx context.Context, msg string, args ...interface{}) {
	newLogger := l.logWithContext(ctx)
	newLogger.Warn(fmt.Sprintf(msg, args...))
}
