package zaplog

import (
	"context"
	"github.com/weflux/x/log"
	"github.com/weflux/x/slices"
	"go.uber.org/zap"
)

var _ log.Logger = new(ZapLogger)

func New(
	zlog *zap.Logger,
) *ZapLogger {
	return &ZapLogger{
		zlog: zlog,
		ctx:  context.Background(),
	}
}

func NewProduction() *ZapLogger {
	zlog, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return New(zlog)
}

func NewDevelopment() *ZapLogger {
	zlog, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	return New(zlog)
}

type ZapLogger struct {
	zlog *zap.Logger
	ctx  context.Context
}

func (z *ZapLogger) WithContext(ctx context.Context) log.Logger {
	return &ZapLogger{
		zlog: z.zlog,
		ctx:  log.MergeContext(z.ctx, ctx),
	}
}

func (z *ZapLogger) WithValues(kvs ...interface{}) log.Logger {
	z.ctx = log.Context(z.ctx, kvs...)
	return &ZapLogger{
		zlog: z.zlog,
		ctx:  z.ctx,
	}
}

func (z *ZapLogger) Fatal(msg string, kvs ...interface{}) {
	var fields []zap.Field
	for k, v := range slices.SliceToMap(kvs) {
		fields = append(fields, zap.Any(k.(string), v))
	}
	z.zlog.Fatal(msg, fields...)
}

func (z *ZapLogger) Error(msg string, err error, kvs ...interface{}) {
	var fields []zap.Field
	fields = append(fields, zap.Any("error", err))
	for k, v := range slices.SliceToMap(kvs) {
		fields = append(fields, zap.Any(k.(string), v))
	}
	z.zlog.Error(msg, fields...)
}

func (z *ZapLogger) Warn(msg string, kvs ...interface{}) {

	var fields []zap.Field
	for k, v := range slices.SliceToMap(kvs) {
		fields = append(fields, zap.Any(k.(string), v))
	}
	z.zlog.Warn(msg, fields...)
}

func (z *ZapLogger) Debug(msg string, kvs ...interface{}) {

	var fields []zap.Field
	for k, v := range slices.SliceToMap(kvs) {
		fields = append(fields, zap.Any(k.(string), v))
	}
	z.zlog.Debug(msg, fields...)
}

func (z *ZapLogger) Info(msg string, kvs ...interface{}) {

	var fields []zap.Field
	for k, v := range slices.SliceToMap(kvs) {
		fields = append(fields, zap.Any(k.(string), v))
	}
	z.zlog.Info(msg, fields...)
}
