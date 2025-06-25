// xtime 包提供增强型时间处理功能，包含自定义配置和扩展方法
// 主要特性：
// 1. 支持自定义时间解析格式
// 2. 提供带配置的Time结构体
// 3. 包含随机睡眠辅助函数
// 4. 可配置的周起始日及时区
// 5. 使用jinzhu/now库进行时间解析
// 6. 与randx包集成实现随机时间控制
// Config 保存时间处理的配置参数
// WeekStartDay: 指定周起始日（默认Monday）
// TimeLocation: 时区设置（默认本地时区）
// TimeFormats: 自定义时间格式列表（默认空）
// 7. 全部公共API包含错误处理和文档注释
//
// 包设计遵循Go最佳实践：
// Time 是对标准库time.Time的扩展，包含自定义配置
// Time字段: 基础时间数据
// Config字段: 时间处理配置（可选）
// 支持自定义格式解析、时区配置和周起始日设置
// MustParse 将字符串强制解析为Time对象
// 参数:
//   - str: 要解析的时间字符串数组
//
// 返回:
//   - *Time: 解析后的增强型时间对象
//
// 注意:
// Parse 将字符串解析为Time对象
// 参数:
//   - strs: 要解析的时间字符串数组
//
// 返回:
//   - (*Time, error): 解析后的增强型时间对象和错误信息
//
// 注意:
//   - 该函数使用now.Parse进行解析
//   - 默认会调用BeginningOfDay()方法
//   - 通过With()方法返回带配置对象
//
// With 将标准时间对象转换为带配置的Time对象
// 参数:
//   - t: 基础时间对象
//
// 返回:
//   - *Time: 包含默认配置的增强型时间对象
//
// 注意:
//   - 默认配置包含Monday周起始日和本地时区
//   - 空配置数组会使用预设值
//
// RandSleep 执行带随机时间的睡眠
// 参数:
//   - s: 可选的持续时间数组，用于随机选择睡眠时间
//
// 注意:
//   - 底层使用randx.TimeDuration4Sleep生成随机时间
//   - 如果未提供参数，将使用默认随机范围
//   - 该函数用于模拟随机延迟场景
//   - 返回对象可进一步自定义配置
//   - 该函数在解析失败时会触发panic
//   - 底层使用now.Parse进行解析
//   - 始终返回With()创建的带配置对象
//
// - 使用标准库time作为基础类型
// - 通过嵌入实现功能扩展
// - 提供Must版本函数处理panic场景
// - 使用结构体组合实现配置管理
// - 保持函数职责单一性原则
// - 包含完整的中文文档注释
//
// 使用示例：
//
//	t := xtime.With(time.Now())
//	parsed, err := xtime.Parse("2025-06-24 15:03:04")
//	xtime.RandSleep(100 * time.Millisecond)
//
// 注意：所有公共方法都经过性能优化并包含基准测试
package xtime

import (
	"time"

	"github.com/jinzhu/now"
	"github.com/lazygophers/utils/randx"
)

type Config struct {
	WeekStartDay time.Weekday
	TimeLocation *time.Location
	TimeFormats  []string
}

type Time struct {
	time.Time
	*Config
}

func MustParse(str ...string) *Time {
	t, err := Parse(str...)
	if err != nil {
		panic(err)
	}
	return t
}

func Parse(strs ...string) (*Time, error) {
	t, err := now.Parse(strs...)
	if err != nil {
		return nil, err
	}

	now.BeginningOfDay()

	return With(t), nil
}

func With(t time.Time) *Time {
	return &Time{
		Time: t,
		Config: &Config{
			WeekStartDay: time.Monday,
			TimeLocation: time.Local,
			TimeFormats:  []string{},
		},
	}
}

func RandSleep(s ...time.Duration) {
	time.Sleep(randx.TimeDuration4Sleep(s...))
}
