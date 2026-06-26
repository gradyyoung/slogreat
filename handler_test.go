package slogreat

import (
	"context"
	"log/slog"
	"os"
	"testing"
)

func TestSlogreatHandler_Handle(t *testing.T) {
	// 初始化测试处理器，配置为无色模式以便于字符串匹配
	h := NewHandler(os.Stdout, &Options{
		NoColor:    false,
		TreeIndent: true,
		AddSource:  true,
	})
	logger := slog.New(h)

	ctx := context.WithValue(context.Background(), "trace_id", "TEST-123")
	logger.InfoContext(ctx, "测试消息", "key", "val")
	logger.Info("测试消息", "key", "val")
}
