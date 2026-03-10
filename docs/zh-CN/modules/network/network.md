---
title: network
---

# network

`network` 提供的是**网络环境与地址相关的辅助能力**，适合处理本机 IP、网卡信息或运行环境里的网络基础信息。

## 适合什么场景

- 启动时探测本机网络信息。
- 需要获取监听地址、出口地址或网卡相关信息。
- 想把一些部署环境相关的网络辅助逻辑收口。

## 使用建议

- 网络信息高度依赖当前机器、容器、网卡和运行环境，示例结果不应当成固定行为。
- 涉及公网 / 内网、多网卡或容器网络时，要在真实部署环境验证。
- 如果目标是 URL 规范化而不是网络探测，请看 [urlx](/modules/network/urlx)。

## 相关文档

- [urlx](/modules/network/urlx)
- [cryptox](/modules/network/cryptox)
- [系统与配置](/modules/system/)
