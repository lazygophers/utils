# 统一 Build Tags 规范

本项目采用统一的 build tags 规范来管理多语言支持。

## 规范说明

### 1. 基础规范

所有语言文件使用以下格式的 build tags：

```go
//go:build i18n_<language> || i18n_all
```

其中 `<language>` 为标准语言代码：

- `en` - 英语（默认，无需 build tag）
- `zh_cn` - 简体中文
- `zh_tw` - 繁体中文  
- `ja` - 日语
- `ko` - 韩语
- `fr` - 法语
- `es` - 西班牙语
- `ar` - 阿拉伯语
- `ru` - 俄语
- `it` - 意大利语
- `pt` - 葡萄牙语
- `de` - 德语

### 2. 构建命令

#### 构建所有语言支持：
```bash
go build -tags="i18n_all" ./...
```

#### 构建特定语言：
```bash
go build -tags="i18n_zh_cn" ./...
go build -tags="i18n_ja" ./...
go build -tags="i18n_fr" ./...
```

#### 构建多种语言：
```bash
go build -tags="i18n_zh_cn,i18n_ja,i18n_fr" ./...
```

#### 默认构建（仅英语）：
```bash
go build ./...
```

### 3. 测试命令

#### 测试所有语言：
```bash
go test -tags="i18n_all" ./...
```

#### 测试特定语言：
```bash
go test -tags="i18n_zh_cn" ./...
```

### 4. 文件命名规范

语言文件应遵循以下命名规范：

- 基础文件：`package_name.go`（英语，无需 build tag）
- 中文简体：`package_name_zh_cn.go`
- 中文繁体：`package_name_zh_tw.go`
- 日语：`package_name_ja.go`
- 韩语：`package_name_ko.go`
- 法语：`package_name_fr.go`
- 西班牙语：`package_name_es.go`
- 阿拉伯语：`package_name_ar.go`
- 俄语：`package_name_ru.go`
- 意大利语：`package_name_it.go`
- 葡萄牙语：`package_name_pt.go`
- 德语：`package_name_de.go`

### 5. 迁移指南

#### 从旧的 build tags 迁移：

**human 包：**
- 旧：`//go:build human_zh || human_all`
- 新：`//go:build i18n_zh_cn || i18n_all`

**validator 包：**
- 旧：`//go:build validator_zh || validator_all`  
- 新：`//go:build i18n_zh_cn || i18n_all`

### 6. 最佳实践

1. **默认支持英语**：所有包都应该默认支持英语，无需 build tag
2. **按需构建**：用户可以根据需要构建特定语言支持
3. **完整测试**：CI/CD 应该测试 `i18n_all` 以确保所有语言正常工作
4. **回退机制**：当找不到特定语言时，应该回退到英语
5. **文档同步**：更新语言文件时，同时更新相关文档

### 7. 环境变量

可以通过环境变量控制默认语言：

```bash
export LANG=zh_CN.UTF-8  # 系统语言
export LC_ALL=zh_CN.UTF-8 # 系统本地化
export UTILS_LOCALE=zh-CN # 项目特定语言设置
```

### 8. Makefile 支持

建议在 Makefile 中添加语言相关的构建目标：

```makefile
# 构建所有语言支持
build-i18n-all:
	go build -tags="i18n_all" ./...

# 构建中文支持
build-zh:
	go build -tags="i18n_zh_cn" ./...

# 测试所有语言
test-i18n-all:
	go test -tags="i18n_all" ./...

# 测试中文
test-zh:
	go test -tags="i18n_zh_cn" ./...
```

这种统一的规范确保了：
- 构建标签的一致性
- 易于理解和维护
- 支持按需构建
- 便于自动化和 CI/CD 集成