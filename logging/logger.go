package logging

import (
	"context"
	"sync"
)

type Logger interface {
	WithContext(ctx context.Context) Logger
	With(args ...interface{}) Logger
	Fatal(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
}

type ctxValues struct {
	fields *sync.Map
}

const ctxLogKey = "_logger"

func MergeContext(target context.Context, source context.Context) context.Context {
	fields := ContextFields(source)
	var args []interface{}
	for k, v := range fields {
		args = append(args, k)
		args = append(args, v)
	}
	if len(args) == 0 {
		return target
	}
	return Context(target, args...)
}

func Context(ctx context.Context, args ...interface{}) context.Context {
	cv, ok := ctx.Value(ctxLogKey).(*ctxValues)
	if !ok {
		cv = &ctxValues{
			fields: &sync.Map{},
		}
		for k, v := range SliceToMap(args) {
			cv.fields.Store(k, v)
		}
		//ctx = context.WithValue(ctx, ctxLogKey, cv)
	} else {
		for k, v := range SliceToMap(args) {
			cv.fields.Store(k, v)
		}
	}

	return context.WithValue(ctx, ctxLogKey, cv)
}

func ContextFields(ctx context.Context) map[interface{}]interface{} {
	res := map[interface{}]interface{}{}
	cv, ok := ctx.Value(ctxLogKey).(*ctxValues)
	if ok {
		cv.fields.Range(func(k, v any) bool {
			res[k] = v
			return true
		})
		return res
	}
	return res
}

func SliceToMap(args []interface{}) map[interface{}]interface{} {
	m := map[interface{}]interface{}{}
	if len(args) == 0 {
		return m
	}
	for i := 1; i <= len(args); i += 2 {
		if i >= len(args) {
			m["empty-key"] = args[i-1]
		} else {
			m[args[i-1]] = args[i]
		}
	}
	return m
}
