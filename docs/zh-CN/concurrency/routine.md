---
    title: routine
    ---

    # routine

    围绕 goroutine、worker pool 和任务调度组织并发执行。

    ## 适用场景

    - 你需要统一管理后台任务和并发 worker。
- 你不希望在业务里到处散落裸 `go` 语句。

    ## 你会接触到什么

    - 通常是并发执行的组织层，而不是具体业务处理层。

    ## 使用建议

    - 如果你还要给外部依赖加保护，可以与 `hystrix` 组合使用。

    ## 相关文档

    - [模块总览](/modules/overview)
    - [API 概览](/api/overview)
