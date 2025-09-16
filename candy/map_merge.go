package candy

import "cmp"

// MergeMap merges two maps, with target values overriding source values
func MergeMap[K cmp.Ordered, V any](source, target map[K]V) map[K]V {
	res := make(map[K]V, len(source))

	// Manual clone implementation
	for k, v := range source {
		res[k] = v
	}

	if len(target) > 0 {
		for k, v := range target {
			res[k] = v
		}
	}

	return res
}

// MergeMapGeneric merges two maps with any comparable key type
func MergeMapGeneric[K comparable, V any](source, target map[K]V) map[K]V {
	if source == nil && target == nil {
		return nil
	}
	if source == nil {
		return CloneMapShallow(target)
	}
	if target == nil {
		return CloneMapShallow(source)
	}

	res := make(map[K]V, len(source)+len(target))

	// Copy source map
	for k, v := range source {
		res[k] = v
	}

	// Override with target map values
	for k, v := range target {
		res[k] = v
	}

	return res
}

// CloneMapShallow creates a shallow copy of a map
func CloneMapShallow[K comparable, V any](m map[K]V) map[K]V {
	if m == nil {
		return nil
	}

	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = v
	}

	return result
}