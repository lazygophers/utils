---
title: 分隔符扫描函数空 sep 入口短路
layer: core
category: arch
keywords: [splitKey,分隔符,死循环,sepLen,空字符串,扫描,key 切分]
source: fix-test-failures
authored-by: skein-memory
created: 1784133088
---

# 分隔符扫描函数空 sep 入口短路

手写分隔符扫描 + 括号状态机的 key 切分函数（如 anyx `splitKey`），入口必须短路：

```go
if len(sep) == 0 { return []string{key} }
```

## 根因

循环内若用 `i += sepLen - 1` 跳过 separator，`sepLen=0` 时 `i += -1` 倒退；且 `key[i:i+sepLen] == sep`（即 `"" == ""`）永真命中 separator 分支 → i 倒退死循环，测试卡死/timeout。

## 模板强制项

凡新写「分隔符扫描 + sepLen 跳转」类函数，入口 `len(sep)==0` guard 作为强制模板项，禁依赖后续 `i+=sepLen-1` 自然处理。

## 关联

- 修复点: anyx/map_any.go `splitKey` + map_any_benchmark_test.go 6 个同模式对照变体（splitKeyCurrent/PreAlloc/...）
- 对比: splitKeyStringsSplit（纯 strings.Split 委托）不经此坑
