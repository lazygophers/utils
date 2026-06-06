---
title: xtime955
---

# xtime955

955 工作制時間常量（9:00-17:00，每週 5 天）。基於標準庫 `time` 包擴展，提供工作日、工作週、工作月等常量。

```go
import "github.com/lazygophers/utils/xtime/xtime955"

// 工作日 = 8 小時
xtime955.WorkDay

// 工作週 = 5 個工作日
xtime955.WorkWeek
```

完整的導出常量參見 [pkg.go.dev](https://pkg.go.dev/github.com/lazygophers/utils/xtime/xtime955)。

相關文檔見 [xtime](/zh-TW/modules/time/xtime)。
