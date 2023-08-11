package main

import (
	"testing"

	"github.com/dwarvesf/test-log/logger"
)

func BenchmarkZeroLog(b *testing.B) {
	log := logger.NewZeroLogger()
	benchmarkLogger(log, b)
}

func BenchmarkSlog(b *testing.B) {
	log := logger.NewSlogLogger()
	benchmarkLogger(log, b)
}

func BenchmarkZapLog(b *testing.B) {
	log := logger.NewZapLogger()
	benchmarkLogger(log, b)
}

func benchmarkLogger(log logger.Log, b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		log.Info("Benchmarking log performance", "iteration", i)
	}
}
