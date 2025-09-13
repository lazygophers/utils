# 开发文档

LazyGophers Utils 项目的开发指南、贡献流程和最佳实践。

## 🚀 快速开始

### 开发环境设置

#### 1. 基础要求
- **Go 版本**: 1.24.0 或更高
- **Git**: 用于版本控制
- **IDE**: 推荐 GoLand, VS Code 或 Vim/Neovim

#### 2. 克隆项目
```bash
git clone https://github.com/lazygophers/utils.git
cd utils
```

#### 3. 安装依赖
```bash
go mod tidy
```

#### 4. 验证环境
```bash
# 运行测试
go test ./...

# 检查代码格式
go fmt ./...

# 静态分析
go vet ./...
```

## 📋 开发流程

### 贡献流程

#### 1. Fork 项目
- 在 GitHub 上 Fork 项目到个人账户
- Clone Fork 后的仓库到本地

#### 2. 创建特性分支
```bash
git checkout -b feature/your-feature-name
# 或
git checkout -b fix/issue-number
```

#### 3. 开发和测试
```bash
# 开发代码
# ...

# 运行测试
go test ./...

# 检查代码覆盖率
go test -cover ./...

# 运行基准测试
go test -bench=. ./...
```

#### 4. 提交变更
```bash
git add .
git commit -m "feat: add your feature description"
```

#### 5. 推送和创建 PR
```bash
git push origin feature/your-feature-name
```
然后在 GitHub 上创建 Pull Request。

### 分支管理策略

#### 主要分支
- **`master`**: 主分支，稳定版本
- **`develop`**: 开发分支，最新功能
- **`feature/*`**: 功能分支
- **`hotfix/*`**: 热修复分支
- **`release/*`**: 发布分支

#### 分支命名规范
```
feature/module-enhancement
fix/issue-123
hotfix/critical-bug
release/v1.2.0
docs/update-readme
```

## 📝 代码规范

### Go 代码风格

#### 1. 命名规范
```go
// ✅ 推荐：清晰的包名
package candy

// ✅ 推荐：描述性的函数名
func ToString(value interface{}) string

// ✅ 推荐：有意义的变量名
var defaultConfig = Config{
    Timeout: 30 * time.Second,
}

// ❌ 避免：缩写和不清晰的命名
func ToStr(v interface{}) string  // 不好
var cfg = Config{}                // 不好
```

#### 2. 函数设计
```go
// ✅ 推荐：单一职责，清晰的参数和返回值
func ParseDate(dateStr string, format string) (time.Time, error) {
    // 实现...
}

// ✅ 推荐：使用泛型提高类型安全
func Must[T any](value T, err error) T {
    if err != nil {
        panic(err)
    }
    return value
}

// ❌ 避免：过多参数
func ComplexFunction(a, b, c, d, e, f string) error  // 不好
```

#### 3. 错误处理
```go
// ✅ 推荐：明确的错误处理
func ProcessData(data []byte) (*Result, error) {
    if len(data) == 0 {
        return nil, errors.New("data is empty")
    }
    
    result, err := parseData(data)
    if err != nil {
        return nil, fmt.Errorf("failed to parse data: %w", err)
    }
    
    return result, nil
}

// ✅ 推荐：使用日志记录错误
func SaveToDatabase(data *Data) error {
    if err := db.Save(data); err != nil {
        log.Error("Failed to save data", log.Error(err))
        return err
    }
    return nil
}
```

#### 4. 文档注释
```go
// ToString 将任意类型的值转换为字符串
// 支持的类型包括：基础类型、切片、映射、结构体等
// 对于不支持的类型，返回空字符串
//
// 示例：
//   ToString(123)     // "123"
//   ToString(3.14)    // "3.14"
//   ToString(true)    // "1"
func ToString(value interface{}) string {
    // 实现...
}
```

### 项目结构规范

#### 1. 包结构
```
utils/
├── app/                 # 应用生命周期管理
├── candy/              # 数据类型转换
├── config/             # 配置管理
├── cryptox/            # 加密工具
├── xtime/              # 时间处理
│   ├── xtime007/       # 007工作制
│   ├── xtime955/       # 955工作制
│   └── xtime996/       # 996工作制
├── docs/               # 项目文档
└── go.mod              # 模块定义
```

#### 2. 文件组织
```go
// 每个包的典型结构
package/
├── main.go             # 主要功能实现
├── helper.go           # 辅助函数
├── types.go            # 类型定义
├── errors.go           # 错误定义
├── main_test.go        # 测试文件
├── examples_test.go    # 示例测试
└── README.md           # 包文档
```

## 🧪 测试规范

### 测试编写指南

#### 1. 单元测试
```go
func TestToString(t *testing.T) {
    tests := []struct {
        name     string
        input    interface{}
        expected string
    }{
        {"int", 123, "123"},
        {"float", 3.14, "3.14"},
        {"bool_true", true, "1"},
        {"bool_false", false, "0"},
        {"nil", nil, ""},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := ToString(tt.input)
            if result != tt.expected {
                t.Errorf("ToString(%v) = %v, expected %v", 
                    tt.input, result, tt.expected)
            }
        })
    }
}
```

#### 2. 基准测试
```go
func BenchmarkToString(b *testing.B) {
    testValues := []interface{}{
        123, 3.14, true, "hello", []int{1, 2, 3},
    }
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        for _, v := range testValues {
            _ = ToString(v)
        }
    }
}
```

#### 3. 示例测试
```go
func ExampleToString() {
    fmt.Println(ToString(123))
    fmt.Println(ToString(3.14))
    fmt.Println(ToString(true))
    // Output:
    // 123
    // 3.14
    // 1
}
```

### 测试覆盖率要求

| 模块类型 | 最低覆盖率 | 目标覆盖率 |
|---------|------------|------------|
| **核心模块** | 90% | 95%+ |
| **工具模块** | 85% | 90%+ |
| **实验性模块** | 80% | 85%+ |

## 🔄 CI/CD 配置

### GitHub Actions 工作流

```yaml
name: CI
on:
  push:
    branches: [ master, develop ]
  pull_request:
    branches: [ master ]

jobs:
  test:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        go-version: [1.24, 1.23]
        
    steps:
    - uses: actions/checkout@v3
    
    - name: Set up Go
      uses: actions/setup-go@v3
      with:
        go-version: ${{ matrix.go-version }}
        
    - name: Cache dependencies
      uses: actions/cache@v3
      with:
        path: ~/go/pkg/mod
        key: ${{ runner.os }}-go-${{ hashFiles('**/go.sum') }}
        
    - name: Install dependencies
      run: go mod download
      
    - name: Run tests
      run: go test -race -coverprofile=coverage.out ./...
      
    - name: Upload coverage
      uses: codecov/codecov-action@v3
      with:
        file: ./coverage.out
        
    - name: Run benchmarks
      run: go test -bench=. -benchmem ./...
      
    - name: Static analysis
      run: |
        go fmt ./...
        go vet ./...
        golangci-lint run
```

## 📦 发布管理

### 版本控制策略

#### 语义版本控制
- **主版本号 (MAJOR)**: 不兼容的 API 更改
- **次版本号 (MINOR)**: 向后兼容的功能增加
- **修订号 (PATCH)**: 向后兼容的错误修复

#### 版本标记格式
```
v1.2.3
v2.0.0-alpha.1
v1.5.0-beta.2
v1.0.0-rc.1
```

### 发布流程

#### 1. 准备发布
```bash
# 更新版本号
git tag v1.2.3

# 生成变更日志
git log --oneline v1.2.2..HEAD > CHANGELOG.md
```

#### 2. 创建发布
```bash
# 推送标签
git push origin v1.2.3

# 在 GitHub 创建 Release
# 包含变更日志和二进制文件
```

## 🔧 开发工具

### 推荐工具和插件

#### IDE 插件
- **VS Code**: Go 扩展包
- **GoLand**: 内置 Go 支持
- **Vim/Neovim**: vim-go 插件

#### 代码质量工具
```bash
# 安装 golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# 安装 goimports
go install golang.org/x/tools/cmd/goimports@latest

# 安装 govulncheck
go install golang.org/x/vuln/cmd/govulncheck@latest
```

#### 性能分析工具
```bash
# pprof
go tool pprof

# trace
go tool trace

# benchcmp
go install golang.org/x/tools/cmd/benchcmp@latest
```

### 开发环境配置

#### .gitignore 配置
```gitignore
# 二进制文件
*.exe
*.exe~
*.dll
*.so
*.dylib

# 测试输出
*.test
*.out
coverage.out

# IDE 文件
.vscode/
.idea/
*.swp
*.swo

# 临时文件
*.tmp
*.log
```

#### pre-commit hooks
```bash
#!/bin/sh
# .git/hooks/pre-commit

# 格式化代码
go fmt ./...

# 静态分析
go vet ./...

# 运行测试
go test -short ./...

# 检查是否有未提交的变更
if ! git diff --exit-code --quiet; then
    echo "Code formatting changed files, please add them and commit again"
    exit 1
fi
```

## 🛡️ 安全指南

### 代码安全

#### 1. 输入验证
```go
// ✅ 推荐：验证输入参数
func ProcessUserInput(input string) error {
    if len(input) == 0 {
        return errors.New("input cannot be empty")
    }
    
    if len(input) > 1000 {
        return errors.New("input too long")
    }
    
    // 进一步处理...
    return nil
}
```

#### 2. 敏感信息处理
```go
// ✅ 推荐：避免在日志中记录敏感信息
func Login(username, password string) error {
    // ❌ 避免：记录密码
    // log.Info("Login attempt", "username", username, "password", password)
    
    // ✅ 推荐：只记录非敏感信息
    log.Info("Login attempt", "username", username)
    
    return nil
}
```

#### 3. 依赖安全
```bash
# 定期检查漏洞
govulncheck ./...

# 更新依赖
go get -u ./...
go mod tidy
```

## 🤝 社区指南

### 交流规范

#### 1. Issue 报告
```markdown
### 问题描述
简明扼要地描述问题

### 复现步骤
1. 执行 `go run main.go`
2. 调用 `candy.ToString(nil)`
3. 观察到错误输出

### 期望行为
应该返回空字符串

### 实际行为
抛出 panic

### 环境信息
- Go 版本: 1.24.0
- 操作系统: macOS 14.0
- 架构: arm64
```

#### 2. Pull Request 模板
```markdown
### 变更描述
简要说明此 PR 的目的和内容

### 变更类型
- [ ] Bug 修复
- [ ] 新功能
- [ ] 性能优化
- [ ] 文档更新
- [ ] 重构

### 测试
- [ ] 新增测试用例
- [ ] 现有测试通过
- [ ] 更新文档

### 检查清单
- [ ] 代码遵循项目规范
- [ ] 测试覆盖率达标
- [ ] 文档已更新
```

## 🔗 相关资源

### 内部文档
- **[API 文档](../api/)** - 完整的 API 参考
- **[模块文档](../modules/)** - 各模块详细说明
- **[测试文档](../testing/)** - 测试策略和报告

### 外部资源
- [Go 官方文档](https://golang.org/doc/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Go 项目布局标准](https://github.com/golang-standards/project-layout)

---

*开发文档最后更新: 2025年09月13日*