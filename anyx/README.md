# anyx
轻量级通用类型转换与数据操作工具

## 功能概述
提供以下核心能力：
- 基本类型（字符串/整数/浮点/布尔）的无损转换
- 复杂结构（切片/映射/指针）的深度遍历与复制
- 支持JSON/YAML的类型安全解析
- 高效的反射操作与数据结构转换

## 核心函数
| 函数名               | 功能描述                     |
|----------------------|------------------------------|
| `ToString(val)`       | 将任意类型安全转为字符串      |
| `ToInt(val)`          | 将任意类型转为整数（错误返回0）|
| `DeepEqual(a, b)`     | 深度比较两个对象是否相等      |
| `Map.Get(key)`        | 支持嵌套路径访问的映射操作    |
| `Slice2Map(list)`     | 将切片转为键值映射            |

## 使用示例
```go
// 基础类型转换
s := anyx.ToString(42) // "42"
i := anyx.ToInt("123") // 123

// 深度复制对象
a := struct{ X int }{42}
b := struct{ X int }{}
anyx.DeepCopy(&a, &b)

// 嵌套路径访问
m := anyx.NewMap(map[string]interface{}{
    "user": map[string]interface{}{
        "id": 123,
    },
})
id := m.GetInt("user.id") // 123
```
