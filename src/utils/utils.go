package utils

func FuncMap[T interface{}, K interface{}](f func(T) K, list []T) []K {
	ret := make([]K, len(list))
	for i, v := range list {
		ret[i] = f(v)
	}
	return ret
}

func All(list []bool) bool {
	for _, v := range list {
		if !v {
			return false
		}
	}
	return true
}

func Contains[T comparable](s []T, item T) bool {
	for _, v := range s {
		if v == item {
			return true
		}
	}
	return false
}

func RemoveIndex[T interface{}](s []T, i int) []T {
	return append(s[:i], s[i+1:]...)
}

func Clone[T interface{}](s []T) []T {
	ret := make([]T, len(s))
	copy(ret, s)
	return ret
}
