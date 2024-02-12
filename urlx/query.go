package urlx

import (
	"github.com/lazygophers/utils/candy"
	"net/url"
)

// SortQuery 对URL参数进行排序
func SortQuery(query url.Values) url.Values {
	if len(query) == 0 {
		return query
	}

	keys := make([]string, 0, len(query))
	for key := range query {
		keys = append(keys, key)
	}

	nq := url.Values{}
	for _, key := range candy.Sort(keys) {
		nq.Set(key, query.Get(key))
	}

	return nq
}
