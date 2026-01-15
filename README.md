# Tokens 价格计算器

一个基于 [Fyne](https://fyne.io/) 的桌面小工具，用于快速换算 **Token 单价** 与 **消费/用量**。

## 功能特性

- **💰 单价设置**：支持输入 `元/百万 tokens` 或 `元/千 tokens`，自动互相换算
- **📊 消费计算**：输入 tokens 数量，计算消费金额（元）
- **🔄 反向计算**：输入消费金额（元），反算可用 tokens 数量（自动显示 `K / M`）
- **📐 单价计算**：输入消费金额和 tokens 数量，反算单价（元/百万 tokens）

## 下载安装

### 直接下载

从 [GitHub Releases](https://github.com/dmxapi/tokens_hs/releases) 下载最新版本的预编译可执行文件，双击运行即可。

### 手动构建

仓库内已提供构建脚本：

- PowerShell：`.\build.ps1`
- CMD：`build.bat`

等价的手动构建命令：

```powershell
$env:CGO_ENABLED = "1"
go build -ldflags "-s -w -H windowsgui" -o tokens_hs.exe
```

> **说明**：`-H windowsgui` 用于生成不弹出控制台窗口的 GUI 程序。构建脚本默认将 `C:\msys64\mingw64\bin` 加入 `PATH`，提供 CGO 所需的编译工具链。

## 使用说明

1. 在「💰 单价设置」里输入任意一个单价（`元/百万 tokens` 或 `元/千 tokens`），另一个会自动换算
2. 选择所需计算方式：
   - 「📊 消费计算」：输入 tokens 数量 → 查看消费金额
   - 「🔄 反向计算」：输入消费金额 → 查看可用 tokens 数量
   - 「📐 单价计算」：输入消费金额和 tokens 数量 → 查看单价

## 项目结构

```
tokens_hs/
├── main.go              # 程序入口
├── calculator.go        # 核心计算逻辑
├── ui/                  # UI 模块
│   ├── app.go           # 应用启动
│   ├── cards.go         # 功能卡片组件
│   ├── components.go    # 通用组件
│   ├── layout.go        # 布局构建
│   ├── state.go         # 状态管理
│   └── utils.go         # 工具函数
├── icon.png             # 应用图标
├── build.ps1            # PowerShell 构建脚本
├── build.bat            # CMD 构建脚本
├── go.mod               # Go 模块定义
└── .github/workflows/   # CI/CD 配置
```

## 技术栈

- **语言**：Go 1.25+
- **GUI 框架**：[Fyne v2](https://fyne.io/) - 跨平台 GUI 框架
- **主题**：暗色主题
- **CI/CD**：GitHub Actions 自动发布

## 相关链接

- [GitHub 仓库](https://github.com/dmxapi/tokens_hs)
- [CNB 仓库](https://cnb.cool/dmxapi/tokens_hs)
