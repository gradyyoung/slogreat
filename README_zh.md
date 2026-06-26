# Slogreat

[![Go Reference](https://pkg.go.dev/badge/github.com/gradyyoung/slogreat.svg)](https://pkg.go.dev/github.com/gradyyoung/slogreat)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**Slogreat** 是一个基于 Go 标准库 `slog` 的高性能、高颜值日志美化处理器。它专为开发环境设计，旨在提供类似 Spring Boot 的日志阅读体验。

### 🌟 核心特性
- **现代化样式**：色彩分明，布局整洁。
- **树状参数显示**：Attr 属性支持 `tree` 命令风格的纵向缩进打印。
- **智能 TraceID 探测**：自动识别 `ctx` 中的链路 ID（支持自定义 Key 或内置常用 Key 搜索）。
- **路径智能缩写**：仿 Spring Boot，将路径如 `/users/project/main.go` 简化为 `u/p/main.go`。
- **高性能**：采用零配置原则，低内存分配，兼容所有 Go 版本。

### 🚀 快速开始

```go
handler := slogreat.NewHandler(os.Stdout, &slogreat.Options{
    Level:      slog.LevelDebug,
    TreeIndent: true,        // 开启树状排版
    AddSource:  true,        // 打印代码路径
    TraceKey:   "trace-key", // 自定义trace-key
})
slog.SetDefault(slog.New(handler))

slog.Info("服务启动成功", "port", 8080, "env", "prod")
```