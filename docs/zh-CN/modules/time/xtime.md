---
title: xtime
---

# xtime

`xtime` 不是普通的“日期格式化小工具”，而是一组围绕**日历、农历、节气和排班时间规则**构建的时间能力。

## 适合什么场景

- 需要从当前时间快速拿到完整日历信息。
- 业务涉及农历、生肖、节气或中文日历表达。
- 想基于仓库内置规则处理固定工作制 / 排班制时间语义。

## 常用入口

- `xtime.NowCalendar()`：获取当前日历对象。
- `(*Calendar).LunarDate()`：读取农历日期。
- `(*Calendar).Animal()`：读取生肖。
- `(*Calendar).CurrentSolarTerm()`：读取当前节气。

## 使用建议

- 这部分规则具有明显领域语义，不要把它当成普通 `time.Time` 封装来理解。
- 涉及节假日、农历或排班制度时，先确认你的业务规则是否真的与包内模型一致。
- 修改前建议先看 `xtime/AGENTS.md`，避免用通用时间直觉误判实现细节。

## 相关文档

- [xtime996](/modules/time/xtime996)
- [xtime955](/modules/time/xtime955)
- [xtime007](/modules/time/xtime007)
