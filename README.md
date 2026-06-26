# Slogreat

[![Go Reference](https://pkg.go.dev/badge/github.com/gradyyoung/slogreat.svg)](https://pkg.go.dev/github.com/gradyyoung/slogreat)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**Slogreat** is a high-performance, aesthetically pleasing log handler for Go's `slog` library. 

*Read this in other languages: [中文版 (Chinese)](README_zh.md)*

### 🌟 Key Features
- **Modern Aesthetic**: Clean, color-coded layouts for better readability.
- **Tree-like Attr Rendering**: Supports `tree`-style indentation for log attributes, making complex data structures easy to parse.
- **Smart TraceID Detection**: Automatically extracts TraceID from `context.Context` (supports both custom keys and common defaults).
- **Intelligent Path Compression**: Simplifies file paths (e.g., `/users/project/main.go` -> `u/p/main.go`) to keep logs concise.
- **High Performance**: Optimized with low memory allocations, compatible with all Go versions (1.21+).

### 🚀 Quick Start

```go
handler := slogreat.NewHandler(os.Stdout, &slogreat.Options{
    Level:      slog.LevelDebug,
    TreeIndent: true,         // Enable tree-style rendering
    AddSource:  true,         // Show file path and line number
    TraceKey:   "trace-key",  // Specify your custom trace key
})
slog.SetDefault(slog.New(handler))

slog.Info("Service started", "port", 8080, "env", "prod")
```