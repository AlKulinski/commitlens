package utils

func DedupBy[T any, K comparable](items []T, keyFn func(T) K) []T {
	seen := make(map[K]bool)
	result := make([]T, 0, len(items))
	for _, item := range items {
		key := keyFn(item)
		if !seen[key] {
			seen[key] = true
			result = append(result, item)
		}
	}
	return result
}
