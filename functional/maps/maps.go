package maps

// ContainsKey returns true if a map contains the given key.
func ContainsKey[M ~map[K]V, K comparable, V any](items M, key K) bool {
	_, ok := items[key]
	return ok
}

// Map calls the function 'f' for each key-value pair and constructs a new map with the key-value
// pairs return by the function.
func Map[M ~map[K]V, K, RK comparable, V, RV any](items M, f func(K, V) (RK, RV)) map[RK]RV {
	result := make(map[RK]RV)
	for k, v := range items {
		retKey, retValue := f(k, v)
		result[retKey] = retValue
	}
	return result
}

// MapKeys calls the function 'f' for each key in the given map and constructs a new map which
// uses the function return value as a key and the original value as the value.
func MapKeys[M ~map[K]V, K, RK comparable, V any](items M, f func(K) RK) map[RK]V {
	result := make(map[RK]V)
	for k, v := range items {
		result[f(k)] = v
	}
	return result
}

// MapValues calls the function 'f' for each value in the given map and constructs a new map
// which uses the original key as the key and the function return value as the value.
func MapValues[M ~map[K]V, K comparable, V, RV any](items M, f func(V) RV) map[K]RV {
	result := make(map[K]RV)
	for k, v := range items {
		result[k] = f(v)
	}
	return result
}

// Merge merges two maps.
// In case of duplicates, the values from the second map have priority.
func Merge[M1, M2 ~map[K]V, K comparable, V any](items1 M1, items2 M2) M1 {
	result := make(M1)
	for k, v := range items1 {
		result[k] = v
	}
	for k, v := range items2 {
		result[k] = v
	}
	return result
}

// MergeFunc merges two maps.
// In case of duplicates, the function 'f' is used to provide the value to be used in the map.
func MergeFunc[M1, M2 ~map[K]V, K comparable, V any](items1 M1, items2 M2, f func(K, V, V) V) M1 {
	result := make(M1)
	for k, v1 := range items1 {
		result[k] = v1
	}
	for k, v2 := range items2 {
		if v1, ok := items1[k]; ok {
			v2 = f(k, v1, v2)
		}
		result[k] = v2
	}
	return result
}

// WithKeys builds a new map from the given map, but using only the specified keys.
func WithKeys[M ~map[K]V, K comparable, V any](items M, keys ...K) M {
	result := make(M)
	for _, key := range keys {
		value, ok := items[key]
		if ok {
			result[key] = value
		}
	}
	return result
}

// WithoutKeys builds a new map from the given map, but skips the specified keys.
func WithoutKeys[M ~map[K]V, K comparable, V any](items M, keys ...K) M {
	result := make(M)
	skip := make(map[K]struct{})
	for _, k := range keys {
		skip[k] = struct{}{}
	}
	for k, v := range items {
		if !ContainsKey(skip, k) {
			result[k] = v
		}
	}
	return result
}
