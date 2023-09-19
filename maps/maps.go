package maps

func MergeMaps(to map[interface{}]interface{}, from map[interface{}]interface{}) map[interface{}]interface{} {
	for k, v := range from {
		to[k] = v
	}
	return to
}
