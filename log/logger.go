package log

import (
	"context"
	"github.com/weflux/x/maps"
	"github.com/weflux/x/slices"
	"strings"
)

type Level int8

const (
	// LevelDebug is logger debug level.
	LevelDebug Level = iota
	// LevelInfo is logger info level.
	LevelInfo
	// LevelWarn is logger warn level.
	LevelWarn
	// LevelError is logger error level.
	LevelError
	// LevelFatal is logger fatal level
	LevelFatal
)

// LevelKey is logger level key.
const LevelKey = "level"

func (l Level) Key() string {
	return LevelKey
}

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return ""
	}
}

// ParseLevel parses a level string into a logger Level value.
func ParseLevel(s string) Level {
	switch strings.ToUpper(s) {
	case "DEBUG":
		return LevelDebug
	case "INFO":
		return LevelInfo
	case "WARN":
		return LevelWarn
	case "ERROR":
		return LevelError
	case "FATAL":
		return LevelFatal
	}
	return LevelInfo
}

type Logger interface {
	WithContext(ctx context.Context) Logger
	WithValues(kvs ...interface{}) Logger
	Fatal(msg string, kvs ...interface{})
	Error(msg string, err error, kvs ...interface{})
	Warn(msg string, kvs ...interface{})
	Debug(msg string, kvs ...interface{})
	Info(msg string, kvs ...interface{})
}

const (
	DefaultMessageKey = "message"
	DefaultErrorKey   = "error"
)

type ValueFunc func(ctx context.Context) interface{}

type ctxValues struct {
	fields map[interface{}]interface{}
}

const ctxLogKey = "_logger"

func MergeContext(target context.Context, source context.Context) context.Context {
	fields := ContextValues(source)
	var kvs []interface{}
	for k, v := range fields {
		kvs = append(kvs, k)
		kvs = append(kvs, v)
	}
	if len(kvs) == 0 {
		return target
	}
	return Context(target, kvs...)
}

func Context(ctx context.Context, kvs ...interface{}) context.Context {
	cv, ok := ctx.Value(ctxLogKey).(*ctxValues)
	if !ok {
		cv = &ctxValues{
			fields: slices.SliceToMap(kvs),
		}
		ctx = context.WithValue(ctx, ctxLogKey, cv)
	} else {
		cv.fields = maps.MergeMaps(cv.fields, slices.SliceToMap(kvs))
	}

	return ctx
}

func ContextValues(ctx context.Context) map[interface{}]interface{} {
	cv, ok := ctx.Value(ctxLogKey).(*ctxValues)
	if ok {
		return cv.fields
	}
	return map[interface{}]interface{}{}
}
