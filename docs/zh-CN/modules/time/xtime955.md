---
title: xtime955
---

# xtime955

955 工作制时间常量（9:00-17:00，每周 5 天）。基于标准库 `time` 包扩展，提供工作日、工作周、工作月等常量。

```go
import "github.com/lazygophers/utils/xtime/xtime955"

// 工作日 = 8 小时
xtime955.WorkDay

// 工作周 = 5 个工作日
xtime955.WorkWeek
```

完整的导出常量参见 [pkg.go.dev](https://pkg.go.dev/github.com/lazygophers/utils/xtime/xtime955)。

相关文档见 [xtime](/modules/time/xtime)。
