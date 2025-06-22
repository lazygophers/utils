# xtime007  
提供基于12小时工作制的时间计算常量  

## 安装指南  
使用以下命令安装包：  
```bash
go get github.com/lazygophers/utils/xtime/xtime007
```

## 使用示例  
### 导入包  
```go
import "github.com/lazygophers/utils/xtime/xtime007"
```

### 时间换算  
```go
// 计算3个完整工作周的总时长
totalWorkTime := xtime007.WorkWeek * 3
fmt.Printf("总工作时长: %s\n", totalWorkTime)
```

### 工时验证  
```go
// 验证是否超过标准工作周
if someDuration > xtime007.WorkWeek {
    fmt.Println("⚠️ 超出标准工作周时长")
}
```

## 贡献说明  
1. **开发环境准备**  
   - 安装Go 1.20+  
   - 设置模块路径：`go mod edit -replace=github.com/lazygophers/utils=../lazygophers/utils`

2. **代码规范**  
   - 新增常量需遵循 `WorkDay`, `WorkWeek` 命名规则  
   - 数学表达式需使用LaTeX格式：`$$1\,\text{WorkWeek} = 5 \times 12\,\text{hours}$$`

3. **提交流程**  
   - 编写测试用例（参考 `xtime/xtime007/xtime_test.go`）  
   - 执行测试：`go test -v ./xtime/xtime007`  
   - 提交PR前运行：`go fmt ./xtime/xtime007`
