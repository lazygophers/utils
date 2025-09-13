# LazyGophers Utils 贡献指南

感谢您对 LazyGophers Utils 的贡献兴趣！本文档为贡献者提供指导和信息。

## 🚀 开始

### 前置要求

- Go 1.24.0 或更高版本
- Git
- Make（可选，用于自动化）

### 开发环境设置

1. **Fork 和 Clone**
   ```bash
   git clone https://github.com/your-username/utils.git
   cd utils
   ```

2. **安装依赖**
   ```bash
   go mod tidy
   ```

3. **验证设置**
   ```bash
   go test ./...
   ```

## 📋 开发指南

### 代码风格

1. **遵循 Go 标准**
   - 使用 `gofmt` 格式化代码
   - 遵循有效的 Go 实践
   - 使用有意义的变量和函数名

2. **包特定指南**
   - 每个包应该独立且可重用
   - 最小化外部依赖
   - 在适当的地方使用泛型以确保类型安全

3. **文档**
   - 所有公共函数必须有中文文档注释
   - 为复杂函数包含使用示例
   - 为关键函数记录性能特征

4. **错误处理**
   - 遵循库的错误处理模式：先记录再返回
   - 使用 `github.com/lazygophers/log` 进行一致的日志记录
   - 提供有意义的错误消息

### 性能指南

1. **内存优化**
   - 为高频操作使用对象池
   - 尽可能使用零拷贝操作
   - 在热路径中最小化内存分配

2. **并发**
   - 尽可能使用原子操作而不是互斥锁
   - 确保并发操作的线程安全
   - 在适当的地方设计无锁算法

3. **基准测试**
   - 为性能关键函数添加基准测试
   - 包含内存分配指标（`-benchmem`）
   - 与基准实现进行比较

## 🧪 测试要求

### 单元测试

1. **测试覆盖率**
   - 新代码的测试覆盖率目标为 90%+
   - 测试成功和错误路径
   - 包含边界情况和边界条件

2. **测试组织**
   ```bash
   # 运行特定包的测试
   go test ./candy
   
   # 运行覆盖率测试
   go test -cover ./...
   
   # 生成覆盖率报告
   go test -coverprofile=coverage.out ./...
   go tool cover -html=coverage.out
   ```

3. **测试命名**
   - 使用描述性测试名称：`TestFunctionName_Condition_ExpectedResult`
   - 使用子测试分组相关测试
   - 为多种场景使用表驱动测试

### 基准测试

1. **性能测试**
   ```bash
   # 运行基准测试
   go test -bench=. -benchmem ./...
   
   # 特定包基准测试
   go test -bench=BenchmarkFunctionName ./package
   ```

2. **基准测试指南**
   - 包含内存分配指标
   - 使用现实的数据大小进行测试
   - 与现有实现进行比较

### 集成测试

1. **跨包测试**
   - 测试包之间的交互
   - 验证与不同 Go 版本的兼容性
   - 测试并发使用模式

## 📝 提交指南

### 提交消息格式

```
<type>(<scope>): <description>

<body>

<footer>
```

**类型：**
- `feat`: 新功能
- `fix`: 错误修复
- `perf`: 性能改进
- `refactor`: 代码重构
- `test`: 添加或更新测试
- `docs`: 文档更改
- `style`: 代码风格更改
- `ci`: CI/CD 更改

**示例：**
```bash
feat(candy): 添加泛型切片转换工具

- 使用 Go 泛型实现 Map、Filter 和 Reduce 函数
- 为所有边界情况添加全面的测试覆盖
- 包含显示 30% 性能提升的基准测试

Closes #123
```

### 分支命名

- 功能分支：`feature/description`
- 错误修复：`fix/description`
- 性能改进：`perf/description`
- 文档：`docs/description`

## 🔍 代码审查流程

### Pull Request 指南

1. **提交前**
   - 确保所有测试通过
   - 运行 `go fmt ./...`
   - 运行 `go vet ./...`
   - 如需要更新文档
   - 为新功能添加或更新测试

2. **PR 描述模板**
   ```markdown
   ## 描述
   更改的简要描述

   ## 更改类型
   - [ ] 错误修复
   - [ ] 新功能
   - [ ] 性能改进
   - [ ] 破坏性更改
   - [ ] 文档更新

   ## 测试
   - [ ] 单元测试通过
   - [ ] 集成测试通过
   - [ ] 包含/更新基准测试
   - [ ] 完成手动测试

   ## 检查清单
   - [ ] 代码遵循风格指南
   - [ ] 完成自我审查
   - [ ] 更新文档
   - [ ] 无破坏性更改（或已记录）
   ```

3. **审查标准**
   - 代码质量和可读性
   - 测试覆盖率和质量
   - 性能影响
   - 破坏性更改
   - 文档完整性

## 🏗️ 架构指南

### 包设计

1. **单一职责**
   - 每个包应该有明确、专注的目的
   - 避免混合不相关的功能
   - 保持公共 API 最小化和清洁

2. **依赖**
   - 最小化外部依赖
   - 尽可能使用标准库
   - 记录依赖的理由

3. **向后兼容性**
   - 在次要版本中避免破坏性更改
   - 在删除前弃用函数
   - 为破坏性更改提供迁移指南

### 性能考虑

1. **内存管理**
   - 为临时对象使用 sync.Pool
   - 在长时间运行的操作中实现适当的清理
   - 在基准测试中监控内存使用

2. **CPU 优化**
   - 分析 CPU 密集型操作
   - 使用适当的数据结构
   - 考虑缓存局部性

3. **并发**
   - 为并发使用设计
   - 有效使用通道和 goroutine
   - 避免竞态条件

## 📚 文档

### API 文档

1. **函数文档**
   ```go
   // ToString 将任意类型转换为字符串
   // 支持基本类型、切片、映射和结构体的转换
   // 对于复杂类型使用JSON序列化
   //
   // 性能特性：
   // - 基本类型转换: O(1)
   // - 复杂类型转换: O(n) where n is serialization complexity
   //
   // 示例：
   //   str := ToString(123)        // "123"
   //   str := ToString([]int{1,2}) // "[1,2]"
   func ToString(v interface{}) string
   ```

2. **包文档**
   - 在包注释中包含包概述
   - 提供使用示例
   - 记录性能特征
   - 包含破坏性更改的迁移指南

### README 指南

1. **包 README**
   - 包目的的清晰描述
   - 安装说明
   - 基本使用示例
   - 性能基准测试
   - API 参考链接

2. **仓库 README**
   - 项目概述
   - 快速开始指南
   - 包目录
   - 贡献指南
   - 许可证信息

## 🚦 发布流程

### 版本管理

1. **语义化版本**
   - MAJOR：破坏性更改
   - MINOR：新功能，向后兼容
   - PATCH：错误修复，向后兼容

2. **发布准备**
   - 更新 CHANGELOG.md
   - 更新相关文件中的版本
   - 确保所有测试通过
   - 更新文档

3. **发布检查清单**
   - [ ] 所有测试通过
   - [ ] 文档已更新
   - [ ] CHANGELOG.md 已更新
   - [ ] 版本已标记
   - [ ] 发布说明已准备

## 🐛 问题指南

### 错误报告

```markdown
**错误描述**
错误的清晰描述

**重现步骤**
1. 第一步
2. 第二步
3. 第三步

**预期行为**
应该发生什么

**实际行为**
实际发生了什么

**环境**
- Go 版本：
- 操作系统：
- 包版本：

**附加上下文**
任何附加信息
```

### 功能请求

```markdown
**功能描述**
提议功能的清晰描述

**用例**
为什么需要这个功能？

**建议实现**
应该如何实现？

**考虑的替代方案**
考虑的其他方法

**附加上下文**
任何附加信息
```

## 🤝 社区指南

### 行为准则

1. **保持尊重**
   - 尊重所有贡献者
   - 在反馈中保持建设性
   - 欢迎新人

2. **合作**
   - 分享知识并帮助他人
   - 提供清晰、有帮助的审查
   - 开放地沟通挑战

3. **专业**
   - 专注于代码，而不是人
   - 优雅地接受批评
   - 给予应得的赞扬

### 沟通渠道

- **GitHub Issues**：错误报告、功能请求
- **GitHub Discussions**：一般问题、想法
- **Pull Requests**：代码审查、讨论

## 📖 学习资源

### Go 最佳实践
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Go Proverbs](https://go-proverbs.github.io/)

### 性能优化
- [Go Performance Tips](https://github.com/golang/go/wiki/Performance)
- [Profiling Go Programs](https://blog.golang.org/profiling-go-programs)

### 测试
- [Testing in Go](https://golang.org/doc/code.html#Testing)
- [Advanced Testing Patterns](https://golang.org/doc/tutorial/add-a-test)

## 📞 获取帮助

如果您需要帮助或有问题：

1. 检查现有文档
2. 搜索现有问题
3. 创建带有清晰描述的新问题
4. 加入我们的社区讨论

感谢您对 LazyGophers Utils 的贡献！您的贡献帮助使这个库对每个人都更好。