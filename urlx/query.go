package urlx

import (
	"net/url"
	"sort"
)

// SortQuery 对URL参数进行排序（零分配优化版本）
func SortQuery(query url.Values) url.Values {
	if len(query) == 0 {
		return query
	}

	// 收集所有key
	keys := make([]string, 0, len(query))
	for key := range query {
		keys = append(keys, key)
	}

	// 原地排序，避免额外分配
	sort.Strings(keys)

	// 预分配结果map的容量
	nq := make(url.Values, len(query))
	for _, key := range keys {
		nq.Set(key, query.Get(key))
	}

	return nq
}
