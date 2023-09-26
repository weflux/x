package zaplog

import (
	"context"
	"github.com/weflux/x/logging"
	"go.uber.org/zap"
)

var _ logging.Logger = new(ZapLogger)

func New(
	zlog *zap.Logger,
) *ZapLogger {
	return &ZapLogger{
		zlog: zlog,
		//ctx:  context.Background(),
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
	//ctx    context.Context
	//fields []zap.Field
}

func (z *ZapLogger) WithContext(ctx context.Context) logging.Logger {
	fields := []zap.Field{}
	for k, v := range logging.ContextFields(ctx) {
		fields = append(fields, zap.Any(k.(string), v))
	}

	return &ZapLogger{
		zlog: z.zlog.With(fields...),
		//fields: fields,
	}
}

func (z *ZapLogger) With(args ...interface{}) logging.Logger {
	//z.ctx = log.Context(z.ctx, args...)
	//fields := z.fields
	fields := []zap.Field{}
	for k, v := range logging.SliceToMap(args) {
		fields = append(fields, zap.Any(k.(string), v))
	}
	return &ZapLogger{
		zlog: z.zlog.With(fields...),
		//ctx:  z.ctx,
		//fields: fields,
	}
}

func (z *ZapLogger) Fatal(msg string, args ...interface{}) {
	//fields := z.fields
	fields := []zap.Field{}
	for k, v := range logging.SliceToMap(args) {
		fields = append(fields, zap.Any(k.(string), v))
	}
	z.zlog.Fatal(msg, fields...)
}

func (z *ZapLogger) Error(msg string, args ...interface{}) {
	//fields := z.fields
	fields := []zap.Field{}
	for k, v := range logging.SliceToMap(args) {
		fields = append(fields, zap.Any(k.(string), v))
	}
	z.zlog.Error(msg, fields...)
}

func (z *ZapLogger) Warn(msg string, args ...interface{}) {

	//fields := z.fields
	fields := []zap.Field{}
	for k, v := range logging.SliceToMap(args) {
		fields = append(fields, zap.Any(k.(string), v))
	}
	z.zlog.Warn(msg, fields...)
}

func (z *ZapLogger) Debug(msg string, args ...interface{}) {

	//fields := z.fields
	fields := []zap.Field{}
	for k, v := range logging.SliceToMap(args) {
		fields = append(fields, zap.Any(k.(string), v))
	}
	z.zlog.Debug(msg, fields...)
}

func (z *ZapLogger) Info(msg string, args ...interface{}) {

	//fields := z.fields
	fields := []zap.Field{}
	for k, v := range logging.SliceToMap(args) {
		fields = append(fields, zap.Any(k.(string), v))
	}
	z.zlog.Info(msg, fields...)
}
