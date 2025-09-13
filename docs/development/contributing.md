# 贡献指南

欢迎为 LazyGophers Utils 项目做出贡献！我们非常感谢社区的每一份力量。

[![Contributors](https://img.shields.io/badge/Contributors-Welcome-brightgreen.svg)](#如何贡献)
[![Code Style](https://img.shields.io/badge/Code%20Style-Go%20Standard-blue.svg)](#代码规范)

## 🤝 如何贡献

### 贡献类型

我们欢迎以下类型的贡献：

- 🐛 **Bug 修复** - 修复已知问题
- ✨ **新功能** - 添加新的工具函数或模块
- 📚 **文档改进** - 完善文档、添加示例
- 🎨 **代码优化** - 性能优化、重构
- 🧪 **测试改进** - 增加测试覆盖率、修复测试问题
- 🌐 **国际化** - 添加多语言支持

### 贡献流程

#### 1. 准备工作

**Fork 项目**
```bash
# 1. Fork 本项目到你的 GitHub 账户
# 2. Clone 你的 fork 到本地
git clone https://github.com/YOUR_USERNAME/utils.git
cd utils

# 3. 添加原项目作为上游仓库
git remote add upstream https://github.com/lazygophers/utils.git

# 4. 创建新的特性分支
git checkout -b feature/your-awesome-feature
```

**设置开发环境**
```bash
# 安装依赖
go mod tidy

# 验证环境
go version  # 需要 Go 1.24.0+
go test ./... # 确保所有测试通过
```

#### 2. 开发阶段

**编写代码**
- 遵循 [代码规范](#代码规范)
- 为新功能编写测试用例
- 确保测试覆盖率不低于现有水平
- 添加必要的文档注释

**提交规范**
```bash
# 使用规范的提交信息格式
git commit -m "feat(module): 添加新的工具函数

- 新增 FormatDuration 函数
- 支持多种时间格式输出
- 添加完整的测试用例
- 更新相关文档

Closes #123"
```

**提交信息格式**：
```
<type>(<scope>): <subject>

<body>

<footer>
```

**Type 类型**：
- `feat`: 新功能
- `fix`: Bug 修复  
- `docs`: 文档更新
- `style`: 代码格式调整
- `refactor`: 代码重构
- `perf`: 性能优化
- `test`: 测试相关
- `chore`: 构建工具或依赖更新

**Scope 范围** (可选)：
- `candy`: candy 模块
- `xtime`: xtime 模块
- `config`: config 模块
- `cryptox`: cryptox 模块
- 等等...

#### 3. 测试验证

**运行测试**
```bash
# 运行所有测试
go test -v ./...

# 检查测试覆盖率
go test -cover -v ./...

# 运行基准测试
go test -bench=. ./...

# 检查代码格式
go fmt ./...

# 静态分析
go vet ./...
```

**性能测试**
```bash
# 运行性能测试
go test -bench=BenchmarkYourFunction -benchmem ./...

# 确保性能没有明显退化
```

#### 4. 创建 Pull Request

**推送到你的 fork**
```bash
git push origin feature/your-awesome-feature
```

**创建 PR**
1. 访问 GitHub 上的项目页面
2. 点击 "New Pull Request"
3. 选择你的分支
4. 填写 PR 描述（参考 [PR 模板](#pr-模板)）
5. 确保通过所有检查

#### 5. 代码审查

- 维护者会审查你的代码
- 根据反馈进行修改
- 保持沟通和合作态度
- 测试通过后将被合并

## 📝 代码规范

### Go 代码风格

**基本规范**
```go
// ✅ 好的示例
package candy

import (
    "context"
    "fmt"
    "time"
    
    "github.com/lazygophers/log"
)

// FormatDuration 格式化时间间隔为人类可读的字符串
// 支持多种精度级别，自动选择合适的单位显示
//
// 参数:
//   - duration: 要格式化的时间间隔
//   - precision: 精度级别 (1-3)
//
// 返回值:
//   - string: 格式化后的字符串，如 "2小时30分钟"
//
// 示例:
//   FormatDuration(90*time.Minute, 2) // 返回 "1小时30分钟"
//   FormatDuration(45*time.Second, 1) // 返回 "45秒"
func FormatDuration(duration time.Duration, precision int) string {
    if duration == 0 {
        return "0秒"
    }
    
    // 实现逻辑...
    return result
}
```

**命名规范**
- 使用 CamelCase（驼峰命名）
- 函数名使用动词开头：`Get`, `Set`, `Format`, `Parse`
- 常量使用全大写：`const MaxRetries = 3`
- 私有成员使用小写开头：`internalHelper`
- 包名使用小写单个单词：`candy`, `xtime`

**注释规范**
- 所有公共函数必须有注释
- 注释以函数名开头
- 包含参数说明、返回值说明  
- 提供使用示例
- 中文注释，简洁明了

**错误处理**
```go
// ✅ 推荐的错误处理方式
func ProcessData(data []byte) (*Result, error) {
    if len(data) == 0 {
        log.Warn("Empty data provided")
        return nil, fmt.Errorf("data cannot be empty")
    }
    
    result, err := parseData(data)
    if err != nil {
        log.Error("Failed to parse data", log.Error(err))
        return nil, fmt.Errorf("parse data failed: %w", err)
    }
    
    return result, nil
}
```

### 项目结构规范

**模块组织**
```
utils/
├── README.md           # 项目总览
├── CONTRIBUTING.md     # 贡献指南  
├── CLAUDE.md          # Claude Code 指令
├── go.mod             # Go 模块定义
├── must.go            # 核心工具函数
├── candy/             # 数据处理工具
│   ├── README.md      # 模块文档
│   ├── to_string.go   # 类型转换
│   └── to_string_test.go
├── xtime/             # 时间处理工具  
│   ├── README.md      # 详细使用文档
│   ├── TESTING.md     # 测试报告
│   ├── PERFORMANCE.md # 性能报告
│   ├── calendar.go    # 日历功能
│   └── calendar_test.go
└── ...
```

**文件命名**
- 使用小写字母和下划线：`to_string.go`
- 测试文件后缀：`_test.go`
- 基准测试：`_benchmark_test.go`
- 文档文件：`README.md`, `TESTING.md`

### 测试规范

**测试覆盖率要求**
- 新功能测试覆盖率必须 ≥ 90%
- 不能降低整体测试覆盖率
- 包含正常用例和边界用例
- 错误处理路径必须测试

**测试示例**
```go
func TestFormatDuration(t *testing.T) {
    testCases := []struct {
        name      string
        duration  time.Duration
        precision int
        want      string
    }{
        {
            name:      "零时间",
            duration:  0,
            precision: 1,
            want:      "0秒",
        },
        {
            name:      "90分钟高精度",
            duration:  90 * time.Minute,
            precision: 2,
            want:      "1小时30分钟",
        },
        // 更多测试用例...
    }
    
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            got := FormatDuration(tc.duration, tc.precision)
            assert.Equal(t, tc.want, got)
        })
    }
}

// 基准测试
func BenchmarkFormatDuration(b *testing.B) {
    duration := 90 * time.Minute
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = FormatDuration(duration, 2)
    }
}
```

## 🎯 开发重点领域

### 高优先级

1. **xtime 模块增强**
   - 农历节气功能完善
   - 性能优化
   - 更多文化特色功能

2. **candy 模块扩展**  
   - 类型转换函数
   - 数据处理工具
   - 性能优化

3. **测试覆盖率提升**
   - 目标：所有模块 > 90%
   - 边界用例补充
   - 性能测试完善

### 中优先级

4. **新工具模块**
   - AI/ML 工具函数
   - 云服务集成工具
   - 微服务工具

5. **文档完善**
   - API 参考文档
   - 最佳实践指南
   - 性能优化指南

### 欢迎贡献的功能

- 🌏 **多语言支持** - 英文文档、错误信息国际化
- 📊 **更多数据格式支持** - XML, YAML, TOML 处理
- 🔧 **开发工具** - 代码生成、配置管理
- 🎨 **UI/UX 工具** - 颜色处理、格式化输出
- 🔐 **安全工具** - 加密解密、签名验证

## 📋 PR 模板

创建 PR 时请使用以下模板：

```markdown
## 变更描述

简要描述本次变更的内容和目的。

## 变更类型

- [ ] Bug 修复
- [ ] 新功能
- [ ] 文档更新
- [ ] 性能优化  
- [ ] 代码重构
- [ ] 测试改进

## 详细变更

### 新增功能
- 新增了 `FormatDuration` 函数
- 支持多种精度级别
- 添加了中文时间单位显示

### 修复问题  
- 修复了时区转换的 bug (#123)
- 解决了内存泄漏问题

### 性能优化
- 优化了字符串拼接性能
- 减少了 30% 的内存分配

## 测试说明

- [ ] 所有测试通过
- [ ] 新增测试用例
- [ ] 测试覆盖率 ≥ 90%
- [ ] 基准测试通过

**测试覆盖率**: 92.5%

## 文档更新

- [ ] 更新了 README.md
- [ ] 添加了函数注释
- [ ] 更新了示例代码

## 兼容性

- [ ] 向后兼容
- [ ] 需要版本升级 (说明原因)
- [ ] 破坏性变更 (详细说明)

## 检查清单

- [ ] 代码遵循项目规范
- [ ] 通过了 `go fmt` 格式检查
- [ ] 通过了 `go vet` 静态检查
- [ ] 所有测试通过
- [ ] 文档已更新
- [ ] 提交信息符合规范

## 相关 Issue

Closes #123
Refs #456

## 截图/演示

如有必要，请提供截图或演示。
```

## 🐛 Bug 报告

发现 Bug？请使用以下模板创建 Issue：

```markdown
## Bug 描述

简要描述遇到的问题。

## 重现步骤

1. 执行步骤 1
2. 执行步骤 2  
3. 观察结果

## 期望行为

描述你期望看到的正确行为。

## 实际行为

描述实际观察到的错误行为。

## 环境信息

- **操作系统**: macOS 12.0
- **Go 版本**: 1.24.0
- **Utils 版本**: v1.2.0
- **其他相关信息**:

## 错误日志

```
paste error logs here
```

## 最小可复现示例

```go
package main

import (
    "github.com/lazygophers/utils/xtime"
)

func main() {
    // 最小的错误复现代码
}
```
```

## ✨ 功能请求

想要新功能？请使用以下模板：

```markdown
## 功能描述

描述你希望添加的功能。

## 使用场景

描述什么情况下会用到这个功能。

## 建议的 API 设计

```go
// 建议的函数签名和使用方式
func NewAwesomeFunction(param string) (Result, error) {
    // ...
}
```

## 替代方案

是否考虑过其他解决方案？

## 额外信息

其他相关信息或参考资料。
```

## 🏆 贡献者认可

### 贡献类型认可

我们会根据贡献类型给予不同的认可：

- 🥇 **核心贡献者** - 长期活跃，重要功能贡献
- 🥈 **积极贡献者** - 多次有价值贡献  
- 🥉 **社区贡献者** - Bug 修复、文档改进
- 🌟 **首次贡献者** - 欢迎第一次贡献

### 贡献统计

我们会在以下地方展示贡献者：

- README.md 贡献者列表
- 发布说明中的致谢
- 项目官网（如有）
- 年度贡献者报告

## 💬 沟通交流

### 获取帮助

- 📖 **文档问题**: 查看各模块的 README.md
- 🐛 **Bug 报告**: [GitHub Issues](https://github.com/lazygophers/utils/issues)
- 💡 **功能讨论**: [GitHub Discussions](https://github.com/lazygophers/utils/discussions)
- ❓ **使用问题**: [GitHub Discussions Q&A](https://github.com/lazygophers/utils/discussions/categories/q-a)

### 讨论规范

请遵循以下沟通规范：

- 使用友善和专业的语言
- 详细描述问题和建议
- 提供足够的上下文信息
- 尊重不同观点和意见
- 积极参与建设性讨论

## 📜 许可证

本项目采用 [GNU Affero General Public License v3.0](LICENSE) 许可证。

**贡献即表示同意**：
- 你拥有所提交代码的版权
- 同意将代码以 AGPL v3.0 许可证发布
- 遵守项目的贡献者行为准则

## 🙏 致谢

感谢所有为 LazyGophers Utils 项目做出贡献的开发者！

**特别感谢**：
- 所有提交 Issue 和 PR 的贡献者
- 提供建议和反馈的社区成员
- 帮助改进文档的志愿者

---

**Happy Coding! 🎉**

有任何问题随时联系维护者团队，我们很乐意帮助你开始贡献之旅！