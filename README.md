# Tokens 价格计算器（tokens_hs）

一个基于 [Fyne](https://fyne.io/) 的桌面小工具，用于快速换算 **Token 单价** 与 **消费/用量**。

## 功能

- **单价设置**：支持输入 `元/百万 tokens` 或 `元/千 tokens`，自动互相换算
- **消费计算**：输入 tokens 数量，计算消费金额（元）
- **反向计算**：输入消费金额（元），反算可用 tokens 数量（自动显示 `K / M`）

## 使用方式

直接运行 `tokens_hs.exe`（如未生成请先按下方步骤构建）。

1. 在「单价设置」里输入任意一个单价（`元/百万 tokens` 或 `元/千 tokens`）。
2. 选择其一：
   - 在「消费计算」里输入 tokens 数量，查看消费金额
   - 在「反向计算」里输入消费金额，查看可用 tokens 数量

## 构建（Windows）

仓库内已提供脚本：

- PowerShell：`.\build.ps1`
- CMD：`build.bat`

等价的手动构建命令为：

```powershell
$env:CGO_ENABLED = "1"
go build -ldflags "-H windowsgui" -o tokens_hs.exe
```

说明：

- `-H windowsgui` 用于生成不弹出控制台窗口的 GUI 程序。
- `build.ps1` / `build.bat` 默认把 `C:\msys64\mingw64\bin` 加入 `PATH`，用于提供 CGO 所需的编译工具链。
