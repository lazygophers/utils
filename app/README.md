# app
应用构建配置与版本信息管理

## 核心功能
1. 根据构建标签自动识别版本类型（Debug/Test/Alpha/Beta/Release）
2. 存储构建时的环境变量（提交哈希、分支、编译时间等）
3. 提供应用元数据访问接口

## 核心变量
```go
var PackageType ReleaseType // 当前构建版本类型
var Organization = "lazygophers" // 开发组织标识
var Name, Version string // 应用名称与版本
```

## 版本类型枚举
```go
type ReleaseType uint8
const (
    Debug ReleaseType = iota // 调试版本
    Test
    Alpha
    Beta
    Release
)
// 支持 ReleaseType.String() 获取对应字符串标识
```

## 使用示例
```go
import "github.com/lazygophers/utils/app"

func main() {
    if app.PackageType == app.Release {
        fmt.Println("生产环境版本:", app.Tag)
    } else if app.PackageType.IsBeta() {
        fmt.Println("测试版本:", app.Branch)
    }
}
```
