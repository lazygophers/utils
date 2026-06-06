---
title: xtime996
---

# xtime996

996 工作制时间常量（9:00-21:00，每周 6 天）。基于标准库 `time` 包扩展，提供工作日、工作周、工作月等常量。

```go
import "github.com/lazygophers/utils/xtime/xtime996"

// 工作日 = 12 小时
xtime996.WorkDay

// 工作周 = 6 个工作日
xtime996.WorkWeek
```

完整的导出常量参见 [pkg.go.dev](https://pkg.go.dev/github.com/lazygophers/utils/xtime/xtime996)。

相关文档见 [xtime](/modules/time/xtime)。
