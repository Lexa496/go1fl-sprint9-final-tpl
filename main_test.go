package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantLen int
	}{
		{name: "нулевой размер", size: 0, wantLen: 0},
		{name: "отрицательный размер", size: -10, wantLen: 0},
		{name: "размер 1", size: 1, wantLen: 1},
		{name: "размер 100", size: 100, wantLen: 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)
			if len(got) != tt.wantLen {
				t.Fatalf("generateRandomElements(%d): длина = %d, ожидалось %d", tt.size, len(got), tt.wantLen)
			}
			for i, v := range got {
				if v <= 0 {
					t.Errorf("элемент %d = %d, ожидалось положительное число", i, v)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want int
	}{
		{name: "nil слайс", in: nil, want: 0},
		{name: "пустой слайс", in: []int{}, want: 0},
		{name: "один элемент", in: []int{42}, want: 42},
		{name: "максимум в начале", in: []int{100, 2, 3}, want: 100},
		{name: "максимум в конце", in: []int{1, 2, 100}, want: 100},
		{name: "максимум в середине", in: []int{1, 100, 2}, want: 100},
		{name: "все элементы равны", in: []int{7, 7, 7}, want: 7},
		{name: "обычный случай", in: []int{1, 9, 3, 7, 5}, want: 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.in); got != tt.want {
				t.Errorf("maximum(%v) = %d, ожидалось %d", tt.in, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want int
	}{
		{name: "nil слайс", in: nil, want: 0},
		{name: "пустой слайс", in: []int{}, want: 0},
		{name: "один элемент", in: []int{5}, want: 5},
		{name: "меньше CHUNKS элементов", in: []int{1, 2, 3}, want: 3},
		{name: "ровно CHUNKS элементов", in: []int{1, 2, 3, 4, 5, 6, 7, 8}, want: 8},
		{name: "больше CHUNKS элементов", in: []int{10, 20, 30, 40, 50, 60, 70, 80, 90}, want: 90},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.in); got != tt.want {
				t.Errorf("maxChunks(%v) = %d, ожидалось %d", tt.in, got, tt.want)
			}
		})
	}
}