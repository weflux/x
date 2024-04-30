package slices

func ToMap(kvs ...interface{}) map[interface{}]interface{} {
	m := map[interface{}]interface{}{}
	for i := 2; i < len(kvs); i += 2 {
		m[kvs[i-2]] = m[kvs[i-1]]
	}
	return m
}
