package models

import (
	"fmt"
	"math/big"
)

type fibonacciService struct {
	FibonacciService
}

type FibonacciService interface {
	MapRecursive(n uint8, mem map[uint8]uint64) uint64
	Iterative(n uint8) uint64
	Recursive(n uint8) uint64
	IterativeBig(n uint32) *big.Int
}

type fibonacci struct{}

func NewFibonacciService() FibonacciService {
	return &fibonacciService{
		FibonacciService: &fibonacci{},
	}
}

// MapRecursive approach of calculating the n-th fibonacci sequence number.
// The map is used to avoid redundant calls when recursively solving the algorithm.
// Time Complexity : O(n)
// Space Complexity : O(n)
func (fm fibonacci) MapRecursive(n uint8, mem map[uint8]uint64) uint64 {
	if n <= 1 {
		return uint64(n)
	}

	if val, ok := mem[n]; ok {
		return val
	}

	val := fm.MapRecursive(n-1, mem) + fm.MapRecursive(n-2, mem)
	mem[n] = val

	return val
}

// Recursive approach of calculating the n-th fibonacci sequence number.
// Time Complexity : O(n)
// Space Complexity : O(n)
func (fm fibonacci) Recursive(n uint8) uint64 {
	if n <= 1 {
		return uint64(n)
	}
	return fm.Recursive(n-1) + fm.Recursive(n-2)
}

// Iterative approach of calculating the n-th fibonacci sequence number.
// Time Complexity : O(n)
// Space Complexity : O(1)
func (fm fibonacci) Iterative(n uint8) uint64 {
	if n <= 1 {
		return uint64(n)
	}

	var first, second uint64 = 0, 1

	for i := uint8(2); i <= n; i++ {
		second, first = first, first+second
	}

	return second + first
}

// IterativeBig approach of calculating the n-th fibonacci sequence number using math/big.Int
// Time Complexity : O(n)
// Space Complexity : O(1)
func (fm fibonacci) IterativeBig(n uint32) *big.Int {
	if n <= 1 {
		return big.NewInt(int64(n))
	}

	var first, second = big.NewInt(0), big.NewInt(1)

	for i := uint32(1); i < n; i++ {
		first.Add(first, second)
		second, first = first, second
	}

	fmt.Printf("%s", second)
	return second
}
