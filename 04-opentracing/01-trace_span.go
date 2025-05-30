package main

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

// 初始化 Jaeger tracer
func initJaeger(service string) (opentracing.Tracer, io.Closer) {
	// 配置 Jaeger
	cfg := jaegercfg.Configuration{
		ServiceName: service,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst, // 采样策略：全量采样
			Param: 1,                       // 1=100%采样
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:          true,                                // 打印日志
			CollectorEndpoint: "http://localhost:14268/api/traces", // Jaeger collector 地址
		},
	}

	// 初始化 tracer
	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}

	fmt.Printf("Jaeger tracer initialized for service: %s\n", service)
	return tracer, closer
}

func main() {
	// 1. 初始化 tracer
	tracer, closer := initJaeger("demo-service")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	// 2. 创建一个 root span (代表整个 trace)
	// 这个 span 没有父 span，所以它是一个 trace 的起点
	rootSpan := tracer.StartSpan("main-operation")
	defer rootSpan.Finish()

	// 打印 root span 的信息
	fmt.Printf("Root Span created. TraceID: %s, SpanID: %s\n",
		rootSpan.Context().(jaeger.SpanContext).TraceID(),
		rootSpan.Context().(jaeger.SpanContext).SpanID())

	// 3. 为 span 添加标签和日志
	rootSpan.SetTag("service.type", "backend")
	rootSpan.SetTag("environment", "dev")
	rootSpan.LogFields(
		log.String("event", "started"),
		log.Int("count", 42),
	)

	// 4. 将 span 存入 context 以便传递
	ctx := opentracing.ContextWithSpan(context.Background(), rootSpan)

	// 5. 模拟业务操作
	fmt.Println("\n--- 开始模拟业务操作 ---")
	callDatabase(ctx)
	callExternalService(ctx)
	doLocalProcessing(ctx)
	fmt.Println("--- 业务操作完成 ---\n")

	// 6. 打印 trace 完成信息
	fmt.Printf("Trace completed. View in Jaeger UI with TraceID: %s\n",
		rootSpan.Context().(jaeger.SpanContext).TraceID())
}

// 模拟数据库调用
func callDatabase(ctx context.Context) {
	// 从 context 中获取父 span 并创建子 span
	span, ctx := opentracing.StartSpanFromContext(ctx, "call-database")
	defer span.Finish()

	// 设置数据库相关的标签
	span.SetTag("db.type", "mysql")
	span.SetTag("db.instance", "users")
	span.SetTag("db.statement", "SELECT * FROM users WHERE id = 123")

	// 打印当前 span 信息
	fmt.Printf("  Database Span. ParentID: %s, SpanID: %s\n",
		span.Context().(jaeger.SpanContext).ParentID(),
		span.Context().(jaeger.SpanContext).SpanID())

	// 模拟数据库操作耗时
	time.Sleep(80 * time.Millisecond)

	// 记录操作结果
	span.LogFields(
		log.String("query.result", "success"),
		log.Int("rows.affected", 1),
	)
}

// 模拟外部服务调用
func callExternalService(ctx context.Context) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "call-external-service")
	defer span.Finish()

	span.SetTag("service.name", "payment-service")
	span.SetTag("http.method", "POST")
	span.SetTag("http.url", "https://api.example.com/payments")

	fmt.Printf("  External Service Span. ParentID: %s, SpanID: %s\n",
		span.Context().(jaeger.SpanContext).ParentID(),
		span.Context().(jaeger.SpanContext).SpanID())

	// 模拟网络延迟
	time.Sleep(120 * time.Millisecond)

	// 模拟错误情况
	span.SetTag("error", true)
	span.LogFields(
		log.String("event", "error"),
		log.String("message", "connection timeout"),
	)
}

// 模拟本地处理
func doLocalProcessing(ctx context.Context) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "local-processing")
	defer span.Finish()

	fmt.Printf("  Local Processing Span. ParentID: %s, SpanID: %s\n",
		span.Context().(jaeger.SpanContext).ParentID(),
		span.Context().(jaeger.SpanContext).SpanID())

	// 模拟处理时间
	time.Sleep(30 * time.Millisecond)

	// 记录处理结果
	span.LogFields(
		log.String("event", "processed"),
		log.Int("items.processed", 5),
	)
}
