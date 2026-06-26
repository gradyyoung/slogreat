package slogreat

// CommonTraceKeys 是业界最常见的 TraceID Key
var CommonTraceKeys = []string{
	// --- 最通用的标准名称 ---
	"trace_id", "traceid", "TraceId", "TraceID",

	// --- RequestID 类 (很多负载均衡和 Web 框架习惯用这个) ---
	"request_id", "requestid", "RequestId", "RequestID",
	"x-request-id", "X-Request-ID",

	// --- 链路追踪组件专用 ---
	"x-trace-id", "X-Trace-ID", // 很多网关层透传的 Header
	"ot-traceid",    // OpenTracing
	"uber-trace-id", // Jaeger 常用
	"x-b3-traceid",  // Zipkin/B3 协议

	// --- 特定框架/中间件 ---
	"x-md-global-trace-id", // Kratos 框架
	"gin.ctx.trace_id",     // Gin 常见自定义名
}
