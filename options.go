package slogreat

import "log/slog"

type Options struct {
	// 1. 准入日志级别
	Level slog.Leveler

	// 2. 是否关闭色彩（默认 false 代表开启彩色，符合本地开发直觉）
	NoColor bool

	// 3. 是否开启 Attr 参数的目录树纵向换行（默认 false 为单行横向排版）
	TreeIndent bool

	// 4. TraceID 打印开关：传入保存在 ctx 中的 key 名字。
	// 如果不为空，会自动去 ctx 捞取并紧跟在 Level 后面打印。
	TraceKey string

	// 5. 是否开启路径缩写（如 /users/project/main.go -> /u/p/main.go）
	AddSource bool
}
