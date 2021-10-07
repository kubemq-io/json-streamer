package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var core = initCore()
var encoderConfig = zapcore.EncoderConfig{
	MessageKey: "message",
	LevelKey:   "severity",
	TimeKey:    "time",
	NameKey:    "name",
	CallerKey:  "caller",

	StacktraceKey:  "stack",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    zapcore.CapitalLevelEncoder,
	EncodeTime:     zapcore.ISO8601TimeEncoder,
	EncodeDuration: zapcore.StringDurationEncoder,
	EncodeCaller:   zapcore.FullCallerEncoder,
	EncodeName:     nil,
}

func initCore() zapcore.Core {

	var w zapcore.WriteSyncer
	std, _, _ := zap.Open("stdout")
	w = zap.CombineWriteSyncers(std)

	enc := zapcore.NewJSONEncoder(encoderConfig)
	return zapcore.NewCore(
		enc,
		zapcore.AddSync(w),
		zapcore.DebugLevel)
}

func LogLevelToZapLevel(value string) zapcore.Level {
	switch value {
	case "debug":
		return zap.DebugLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}

type Logger struct {
	*zap.SugaredLogger
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l.Infof(format, v)
}

func NewLogger(name string) *Logger {

	zapLogger := zap.New(core, zap.Hooks(func(entry zapcore.Entry) error {
		if entry.Level < zap.ErrorLevel {
			return nil
		}

		return nil
	}))
	l := &Logger{
		SugaredLogger: zapLogger.Sugar().With("source", name),
	}
	return l
}

func (l *Logger) Noticef(format string, v ...interface{}) {
	str := fmt.Sprintf(format, v...)
	l.Info(str)
}

func (l *Logger) Tracef(format string, v ...interface{}) {
	str := fmt.Sprintf(format, v...)
	l.Debug(str)
}
func (l *Logger) Fatalf(format string, v ...interface{}) {
	str := fmt.Sprintf(format, v...)
	l.DPanicf(str)
}

func (l *Logger) NewWith(p1, p2 string) *Logger {
	return &Logger{SugaredLogger: l.SugaredLogger.With(p1, p2)}
}

func (l *Logger) Write(p []byte) (n int, err error) {
	l.Info(string(p))
	return len(p), nil
}
