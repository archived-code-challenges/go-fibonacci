package models

import (
	"fmt"
	"testing"
)

func BenchmarkFibonacciIterative10(b *testing.B) {
	fs := NewFibonacciService()
	for i := 0; i < b.N; i++ {
		fs.Iterative(10)
	}
}

func BenchmarkFibonacciMapRecursive10(b *testing.B) {
	fs := NewFibonacciService()
	mem := make(map[uint8]uint64)
	for i := 0; i < b.N; i++ {
		fs.MapRecursive(10, mem)
	}
}

func BenchmarkFibonacciRecursive10(b *testing.B) {
	fs := NewFibonacciService()
	for i := 0; i < b.N; i++ {
		fs.Recursive(10)
	}
}

func BenchmarkFibonacciIterativeBig10(b *testing.B) {
	fs := NewFibonacciService()
	for i := 0; i < b.N; i++ {
		fs.IterativeBig(10)
	}
}

func BenchmarkFibonacciIterative20(b *testing.B) {
	fs := NewFibonacciService()
	for i := 0; i < b.N; i++ {
		fs.Iterative(20)
	}
}

func BenchmarkFibonacciMapRecursive20(b *testing.B) {
	fs := NewFibonacciService()
	mem := make(map[uint8]uint64)
	for i := 0; i < b.N; i++ {
		fs.MapRecursive(20, mem)
	}
}

func BenchmarkFibonacciRecursive20(b *testing.B) {
	fs := NewFibonacciService()
	for i := 0; i < b.N; i++ {
		fs.Recursive(20)
	}
}

func TestFibonacciIterative(t *testing.T) {
	fs := NewFibonacciService()

	tests := []struct {
		n    uint8
		want uint64
	}{
		{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {10, 55}, {42, 267914296},
	}

	for _, tt := range tests {
		if got := fs.Iterative(tt.n); got != tt.want {
			t.Errorf("Invalid Fibonacci value for N: %d, got: %d, want: %d", tt.n, got, tt.want)
		}
	}
}

func TestFibonacciMapRecursive(t *testing.T) {
	fs := NewFibonacciService()

	tests := []struct {
		n    uint8
		want uint64
	}{
		{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {10, 55}, {42, 267914296},
	}

	for _, tt := range tests {
		mem := make(map[uint8]uint64)
		if got := fs.MapRecursive(tt.n, mem); got != tt.want {
			t.Errorf("Invalid Fibonacci value for N: %d, got: %d, want: %d", tt.n, got, tt.want)
		}
	}
}

func TestFibonacciRecursive(t *testing.T) {
	fs := NewFibonacciService()

	tests := []struct {
		n    uint8
		want uint64
	}{
		{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {10, 55}, {42, 267914296},
	}

	for _, tt := range tests {
		if got := fs.Recursive(tt.n); got != tt.want {
			t.Errorf("Invalid Fibonacci value for N: %d, got: %d, want: %d", tt.n, got, tt.want)
		}
	}
}

func TestFibonacciIterativeBig(t *testing.T) {
	fs := NewFibonacciService()

	tests := []struct {
		n    uint32
		want string
	}{
		{0, "0"}, {1, "1"}, {2, "1"}, {3, "2"}, {100, "354224848179261915075"}, {400, "176023680645013966468226945392411250770384383304492191886725992896575345044216019675"},
	}

	for _, tt := range tests {
		if got := fs.IterativeBig(tt.n); fmt.Sprintf("%v", got) != tt.want {
			t.Errorf("Invalid Fibonacci value for N: %d, got: %d, want: %s", tt.n, got, tt.want)
		}
	}
}
