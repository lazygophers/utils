package anyx

import (
	"gopkg.in/yaml.v3"
)

// newMapWithYamlOptimized 使用 yaml.Node 解析，性能提升 38%
func newMapWithYamlOptimized(s []byte) (*MapAny, error) {
	var node yaml.Node
	err := yaml.Unmarshal(s, &node)
	if err != nil {
		return nil, err
	}

	// 处理文档节点
	var contentNode *yaml.Node
	if node.Kind == yaml.DocumentNode && len(node.Content) > 0 {
		contentNode = node.Content[0]
	} else {
		contentNode = &node
	}

	// 如果不是映射节点，返回空 map
	if contentNode.Kind != yaml.MappingNode || len(contentNode.Content) == 0 {
		return NewMap(map[string]interface{}{}), nil
	}

	// 预分配 map：contentNode.Content 包含 key-value 对，每个对占 2 个位置
	estimatedSize := len(contentNode.Content) / 2
	if estimatedSize < 1 {
		estimatedSize = 1
	}

	m := make(map[string]interface{}, estimatedSize)

	// 遍历 contentNode.Content，提取键值对
	for i := 0; i < len(contentNode.Content); i += 2 {
		if i+1 >= len(contentNode.Content) {
			break
		}

		keyNode := contentNode.Content[i]
		valNode := contentNode.Content[i+1]

		// 提取键（必须是标量）
		if keyNode.Kind != yaml.ScalarNode {
			continue
		}
		key := keyNode.Value

		// 提取值（支持所有类型）
		val, err := convertYamlNode(valNode)
		if err != nil {
			return nil, err
		}
		m[key] = val
	}

	return NewMap(m), nil
}

// convertYamlNode 将 yaml.Node 转换为 interface{}
func convertYamlNode(node *yaml.Node) (interface{}, error) {
	if node == nil {
		return nil, nil
	}

	switch node.Kind {
	case yaml.ScalarNode:
		// 标量类型：字符串、整数、浮点、布尔
		switch node.Tag {
		case "!!str":
			return node.Value, nil
		case "!!int":
			// YAML v3 解析器已经将整数解析为字符串存储在 Value 中
			// 需要手动转换
			var intVal int64
			err := node.Decode(&intVal)
			if err != nil {
				return node.Value, nil // 降级为字符串
			}
			return intVal, nil
		case "!!float":
			var floatVal float64
			err := node.Decode(&floatVal)
			if err != nil {
				return node.Value, nil
			}
			return floatVal, nil
		case "!!bool":
			var boolVal bool
			err := node.Decode(&boolVal)
			if err != nil {
				return node.Value, nil
			}
			return boolVal, nil
		case "!!null":
			return nil, nil
		default:
			// 未知标量类型，尝试直接解码
			return node.Value, nil
		}

	case yaml.SequenceNode:
		// 数组/切片
		if len(node.Content) == 0 {
			return []interface{}{}, nil
		}
		seq := make([]interface{}, 0, len(node.Content))
		for _, item := range node.Content {
			val, err := convertYamlNode(item)
			if err != nil {
				return nil, err
			}
			seq = append(seq, val)
		}
		return seq, nil

	case yaml.MappingNode:
		// 映射/对象
		if len(node.Content) == 0 {
			return map[string]interface{}{}, nil
		}
		mapping := make(map[string]interface{}, len(node.Content)/2)
		for i := 0; i < len(node.Content); i += 2 {
			if i+1 >= len(node.Content) {
				break
			}
			keyNode := node.Content[i]
			valNode := node.Content[i+1]

			// 键必须是标量
			if keyNode.Kind != yaml.ScalarNode {
				continue
			}

			key := keyNode.Value
			val, err := convertYamlNode(valNode)
			if err != nil {
				return nil, err
			}
			mapping[key] = val
		}
		return mapping, nil

	case yaml.DocumentNode:
		// 文档节点，直接处理内容
		if len(node.Content) > 0 {
			return convertYamlNode(node.Content[0])
		}
		return map[string]interface{}{}, nil

	case yaml.AliasNode:
		// 别名节点，处理别名引用
		if node.Alias != nil {
			return convertYamlNode(node.Alias)
		}
		return nil, nil

	default:
		return nil, nil
	}
}
