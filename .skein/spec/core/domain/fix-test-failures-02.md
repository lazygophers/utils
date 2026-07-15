---
title: 历史日期算整天禁用 time.Local（LMT 偏移）
layer: core
category: domain
keywords: [xtime,农历,lunar,time.Local,LMT,时区,1900,off-by-one,历史日期]
source: fix-test-failures
authored-by: skein-memory
created: 1784133093
---

# 历史日期算整天禁用 time.Local

农历/历史日期（< 1900s 本地时区定义点）做 Unix 秒差算整天时，**禁用 `time.Local`**。

## 根因

`time.Local` 在远古年份（如 1900）是 LMT（Local Mean Time，如 +06:55）而非标准时区（CST +08:00）。日期分量用 Local 解析、基准用 UTC，混算致 Unix 秒差非整天，`/86400` 截断后 offset 偏 1 天。

## 正确做法

日期分量取本地语义（年月日），但做差时用 UTC 重建：

```go
t := time.Unix(ts, 0).In(time.Local)
t1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
t2 := time.Date(1900, 1, 31, 0, 0, 0, 0, time.UTC)
offset = (t1.Unix() - t2.Unix()) / 86400
```

## 警惕

LMT→标准时区切换在各地不同（上海 1901 年）。任何「历史日期做整天差」先问时区来源。
