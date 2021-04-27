package utils

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortWorstCase(t *testing.T) {
	els := []int{6, 5, 4, 3, 2, 1}
	actcual := BubbleSort(els)

	assert.NotNil(t, actcual)

	expected := []int{1, 2, 3, 4, 5, 6}
	assert.Equal(t, expected, actcual)
}

func TestBubbleSortBestCase(t *testing.T) {
	els := []int{1, 2, 3, 4, 5, 6, 7}
	actcual := BubbleSort(els)

	assert.NotNil(t, actcual)

	expected := []int{1, 2, 3, 4, 5, 6, 7}
	assert.Equal(t, expected, actcual)
}

func getElements(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}

func TestGetElements10(t *testing.T) {
	els := getElements(10)
	assert.Equal(t, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, els)
}

func BenchmarkBubbleSort10(b *testing.B) {
	els := getElements(10)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkBubbleSort10000(b *testing.B) {
	els := getElements(10000)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkBubbleSort50000(b *testing.B) {
	els := getElements(50000)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkBubbleSort100000(b *testing.B) {
	els := getElements(100000)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkNativeSort10(b *testing.B) {
	els := getElements(10)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}

func BenchmarkNativeSort10000(b *testing.B) {
	els := getElements(10000)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}

func BenchmarkNativeSort50000(b *testing.B) {
	els := getElements(50000)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}

func BenchmarkNativeSort100000(b *testing.B) {
	els := getElements(100000)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}
