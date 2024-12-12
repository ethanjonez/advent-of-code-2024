package maphelper

func GetValues[V comparable, K comparable](mappy map[K]V) (values []V) {
	for _, v := range mappy {
		values = append(values, v)
	}

	return values
}

func GetKeys[V comparable, K comparable](mappy map[K]V) (keys []K) {
	for k, _ := range mappy {
		keys = append(keys, k)
	}

	return keys
}
