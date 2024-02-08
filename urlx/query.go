package urlx

import (
	"github.com/elliotchance/pie/v2"
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
	for _, key := range pie.Sort(keys) {
		nq.Set(key, query.Get(key))
	}

	return nq
}
