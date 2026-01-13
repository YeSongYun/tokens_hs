# v1.0.0（初版）

## 更新内容

- 提供基于 Fyne 的桌面 GUI「Tokens 价格计算器」
- 支持 `元/百万 tokens` 与 `元/千 tokens` 双向换算
- 支持按 tokens 数量计算消费金额（元）
- 支持按消费金额（元）反算可用 tokens 数量，并以 `K / M` 形式展示
- 优化 tokens 显示格式：去除多余尾零

## 下载

- Windows：在 Release 的 Assets 中下载 `tokens_hs.exe`

## 使用说明

- 先在「单价设置」里输入任意一个单价（`元/百万 tokens` 或 `元/千 tokens`），再进行计算。

## 构建（Windows）

- PowerShell：`.\build.ps1`
- CMD：`build.bat`

