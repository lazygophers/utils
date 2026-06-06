---
title: xtime007
---

# xtime007

全天候工作制时间常量（0:00-24:00，每周 7 天）。基于标准库 `time` 包扩展，提供工作日、工作周、工作月等常量。

```go
import "github.com/lazygophers/utils/xtime/xtime007"

// 工作日 = 24 小时
xtime007.WorkDay

// 工作周 = 7 天
xtime007.WorkWeek
```

完整的导出常量参见 [pkg.go.dev](https://pkg.go.dev/github.com/lazygophers/utils/xtime/xtime007)。

相关文档见 [xtime](/modules/time/xtime)。
