// Package app 提供应用程序环境类型和构建信息管理。
//
// 支持的构建环境：
//   - Debug: 调试模式（默认）
//   - Test: 测试/金丝雀环境
//   - Alpha: 内部测试版本
//   - Beta: 公开测试版本
//   - Release: 正式发布版本
//
// 使用构建标签（-tags）在编译时设置环境：
//	go build -tags=debug
//	go build -tags=test
//	go build -tags=alpha
//	go build -tags=beta
//	go build -tags=release
//
// 环境变量覆盖：
// 可通过 APP_ENV 环境变量在运行时覆盖构建标签设置：
//	APP_ENV=dev|development    → Debug
//	APP_ENV=test|canary        → Test
//	APP_ENV=prod|production    → Release
//	APP_ENV=alpha              → Alpha
//	APP_ENV=beta               → Beta
//
// 示例：
//
//	import "github.com/lazygophers/utils/app"
//
//	if app.Env == app.Debug {
//	    // 开启调试功能
//	}
//
//	fmt.Printf("Running in %s mode\n", app.Env)
package app
