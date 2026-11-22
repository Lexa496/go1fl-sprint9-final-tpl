package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	// SIZE — количество элементов в генерируемом слайсе.
	SIZE = 10_000_000
	// CHUNKS — количество частей, на которые делится слайс.
	CHUNKS = 8
)

// generateRandomElements создаёт слайс положительных случайных чисел заданного размера.
// Для неположительного размера возвращает пустой слайс.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	elements := make([]int, size)
	for i := range elements {
		elements[i] = rand.Intn(1_000_000) + 1
	}
	return elements
}

// maximum возвращает максимальный элемент слайса.
// Для пустого (или nil) слайса возвращает 0.
func maximum(elements []int) int {
	if len(elements) == 0 {
		return 0
	}

	maxVal := elements[0]
	for _, v := range elements[1:] {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

// maxChunks находит максимум, разбивая слайс на CHUNKS частей
// и обрабатывая каждую часть в отдельной горутине.
func maxChunks(elements []int) int {
	if len(elements) == 0 {
		return 0
	}

	chunkSize := len(elements) / CHUNKS
	// Если элементов меньше, чем частей, многопоточность не даст выигрыша.
	if chunkSize == 0 {
		return maximum(elements)
	}

	maxes := make([]int, CHUNKS)
	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		// Последняя часть забирает остаток элементов.
		if i == CHUNKS-1 {
			end = len(elements)
		}

		wg.Add(1)
		go func(chunk []int, idx int) {
			defer wg.Done()
			maxes[idx] = maximum(chunk)
		}(elements[start:end], i)
	}

	wg.Wait()

	return maximum(maxes)
}

func main() {
	elements := generateRandomElements(SIZE)

	start := time.Now()
	maxSingle := maximum(elements)
	singleDuration := time.Since(start).Microseconds()

	start = time.Now()
	maxConcurrent := maxChunks(elements)
	concurrentDuration := time.Since(start).Microseconds()

	fmt.Printf("Однопоточный поиск:   максимум = %d, время = %d мкс\n", maxSingle, singleDuration)
	fmt.Printf("Многопоточный поиск:  максимум = %d, время = %d мкс\n", maxConcurrent, concurrentDuration)
}