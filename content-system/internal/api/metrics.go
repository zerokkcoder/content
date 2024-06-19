package api

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

// prometheus监控中间件 localhost:9090/metrics
func prometheusMiddleware() gin.HandlerFunc {
	// 计数器指标
	requestsTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_request_total",
		Help: "Total number of http request",
	}, []string{
		"method",
		"path",
	})

	// 错误码统计
	requestsCodeTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_request_code_total",
		Help: "Total number of http request code",
	}, []string{
		"method",
		"path",
		"code",
	})

	// 概要指标
	requestDuration := prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: "http_request_duration_seconds",
		Help: "Duration of HTTP requests",
		// 分位数
		Objectives: map[float64]float64{
			0.5:  0.05,
			0.9:  0.01,
			0.99: 0.001,
		},
	}, []string{
		"method",
		"path",
	})

	prometheus.MustRegister(requestsTotal)
	prometheus.MustRegister(requestsCodeTotal)
	prometheus.MustRegister(requestDuration)

	return func(c *gin.Context) {
		start := time.Now()
		method := c.Request.Method
		path := c.FullPath()
		requestsTotal.WithLabelValues(method, path).Inc()

		c.Next()
		elapsed := time.Since(start).Seconds()
		requestDuration.WithLabelValues(method, path).Observe(elapsed)

		code := c.Writer.Status()
		requestsCodeTotal.WithLabelValues(method, path, strconv.Itoa(code)).Inc()
	}
}
