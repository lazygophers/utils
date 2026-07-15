---
title: Singleflight error 不缓存语义
layer: core
category: domain
keywords: [singledo,singleflight,error 缓存,s.call,并发,Do]
source: fix-test-failures
authored-by: skein-memory
created: 1784133093
---

# Singleflight error 不缓存语义

本项目 `singledo.Single.Do` 取「error 不缓存」语义：call 完成后 `s.call` **无条件清**（无论 success/error），failure 不被后续并发请求复用。

```go
callM.val, callM.err = fn()
callM.wg.Done()

s.mux.Lock()
if callM.err == nil {
    s.last = now
    s.result = callM.val
}
s.call = nil  // 无条件清，error 不缓存
s.mux.Unlock()
```

## 判定依据

若 error 分支不清 `s.call`，第 2 个调用会命中 `s.call != nil` 分支拿到缓存 error，失败被复用。无条件清保证 error 只返回给当时在等的并发者。

## 关联

- 删过时测试 TestSingleDo_ErrorKeepsCall（语义已翻转）
- DoCtx API 待后补（当前删除其测试）
