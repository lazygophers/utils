package anyx

import (
	"strconv"
	"strings"
)

// getNested 是基准测试辅助方法，模拟嵌套查找但无defer
func (p *MapAny) getNested(key string) (interface{}, bool) {
	seq := p.seq.Load().(string)
	keys := strings.Split(key, seq)

	if len(keys) <= 1 {
		return nil, false
	}

	currentMap := p.data
	for i := 0; i < len(keys)-1; i++ {
		val, ok := currentMap[keys[i]]
		if !ok {
			return nil, false
		}

		nestedMap, ok := val.(map[string]interface{})
		if !ok {
			mapAny := p.toMap(val)
			if mapAny == nil {
				return nil, false
			}
			nestedMap = mapAny.data
		}
		currentMap = nestedMap
	}

	finalKey := keys[len(keys)-1]
	val, ok := currentMap[finalKey]
	if ok {
		return val, true
	}

	return nil, false
}

// strconvParseUint 是 strconv.ParseUint 的内联版本，用于基准测试
func strconvParseUint(s string, base int, bitSize int) (uint64, error) {
	return strconv.ParseUint(s, base, bitSize)
}
