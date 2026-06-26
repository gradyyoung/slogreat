package slogreat

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

// ANSI 彩色转义码
const (
	colorReset  = "\033[0m"
	colorGray   = "\033[90m"
	colorBlue   = "\033[34m"
	colorCyan   = "\033[36m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorRed    = "\033[31m"
)

// 常见 TraceID 探测列表
var commonTraceKeys = []string{
	"trace_id", "traceid", "TraceId", "TraceID",
	"request_id", "requestid", "RequestId", "RequestID",
	"x-request-id", "X-Request-ID", "x-trace-id", "X-Trace-ID",
}

type SlogreatHandler struct {
	w    io.Writer
	opts Options
	mu   sync.Mutex
}

func NewHandler(w io.Writer, opts *Options) *SlogreatHandler {
	h := &SlogreatHandler{w: w}
	if opts != nil {
		h.opts = *opts
	}
	return h
}

func (h *SlogreatHandler) Enabled(_ context.Context, l slog.Level) bool {
	minLevel := slog.LevelInfo
	if h.opts.Level != nil {
		minLevel = h.opts.Level.Level()
	}
	return l >= minLevel
}

func (h *SlogreatHandler) Handle(ctx context.Context, r slog.Record) error {
	buf := bytes.NewBuffer(make([]byte, 0, 1024))

	// 1. 时间戳
	if !r.Time.IsZero() {
		writeColor(buf, colorGray, !h.opts.NoColor)
		buf.WriteString(r.Time.Format("2006-01-02 15:04:05.000"))
		writeColor(buf, colorReset, !h.opts.NoColor)
		buf.WriteByte(' ')
	}

	// 2. 日志级别
	levelStr := fmt.Sprintf("%-5s", r.Level.String())
	if !h.opts.NoColor {
		color := colorGreen
		switch r.Level {
		case slog.LevelDebug:
			color = colorBlue
		case slog.LevelWarn:
			color = colorYellow
		case slog.LevelError:
			color = colorRed
		}
		buf.WriteString(color + levelStr + colorReset)
	} else {
		buf.WriteString(levelStr)
	}
	buf.WriteByte(' ')

	// 3. 智能 TraceID
	if traceID := findTraceID(ctx, h.opts.TraceKey); traceID != "" {
		writeColor(buf, colorCyan, !h.opts.NoColor)
		buf.WriteString("[" + traceID + "]")
		writeColor(buf, colorReset, !h.opts.NoColor)
		buf.WriteByte(' ')
	}

	// 4. 路径 (兼容所有 Go 版本)
	if h.opts.AddSource && r.PC != 0 {
		fs, _ := runtime.CallersFrames([]uintptr{r.PC}).Next()
		path := compressPath(fs.File)
		writeColor(buf, colorBlue, !h.opts.NoColor)
		buf.WriteString("[" + path + ":" + fmt.Sprint(fs.Line) + "]")
		writeColor(buf, colorReset, !h.opts.NoColor)
		buf.WriteByte(' ')
	}

	buf.WriteString(": ")
	buf.WriteString(r.Message)

	// 5. Attr 渲染
	var attrs []slog.Attr
	r.Attrs(func(a slog.Attr) bool { attrs = append(attrs, a); return true })

	if len(attrs) > 0 {
		if h.opts.TreeIndent {
			buf.WriteByte('\n')
			for i, attr := range attrs {
				if i == len(attrs)-1 {
					buf.WriteString("└── ")
				} else {
					buf.WriteString("├── ")
				}
				appendAttr(buf, attr, h.opts.NoColor)
				if i != len(attrs)-1 {
					buf.WriteByte('\n')
				}
			}
		} else {
			buf.WriteString("  ")
			for _, attr := range attrs {
				appendAttr(buf, attr, h.opts.NoColor)
				buf.WriteByte(' ')
			}
		}
	}
	buf.WriteByte('\n')

	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := h.w.Write(buf.Bytes())
	return err
}

// 辅助工具函数
func findTraceID(ctx context.Context, userKey string) string {
	if userKey != "" {
		if val := ctx.Value(userKey); val != nil {
			if s, ok := val.(string); ok && s != "" {
				return s
			}
		}
	}
	for _, k := range commonTraceKeys {
		if val := ctx.Value(k); val != nil {
			if s, ok := val.(string); ok && s != "" {
				return s
			}
		}
	}
	return ""
}

func writeColor(buf *bytes.Buffer, code string, active bool) {
	if active {
		buf.WriteString(code)
	}
}

func compressPath(path string) string {
	dir, file := filepath.Split(path)
	if dir == "" {
		return file
	}
	parts := strings.Split(strings.Trim(dir, string(filepath.Separator)), string(filepath.Separator))
	for i, part := range parts {
		if len(part) > 0 {
			parts[i] = string(part[0])
		}
	}
	return filepath.Join(strings.Join(parts, string(filepath.Separator)), file)
}

func appendAttr(buf *bytes.Buffer, attr slog.Attr, noColor bool) {
	if !noColor {
		buf.WriteString(colorCyan + attr.Key + colorReset + "=" + colorGreen + fmt.Sprint(attr.Value.Any()) + colorReset)
	} else {
		buf.WriteString(attr.Key + "=" + fmt.Sprint(attr.Value.Any()))
	}
}

func (h *SlogreatHandler) WithAttrs(attrs []slog.Attr) slog.Handler { return h }
func (h *SlogreatHandler) WithGroup(name string) slog.Handler       { return h }
