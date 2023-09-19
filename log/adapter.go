package log

type Adapter interface {
	Log(level Level, kvs ...interface{})
}
