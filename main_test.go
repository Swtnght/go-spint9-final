package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	t.Run("нулевой размер даёт nil", func(t *testing.T) {
		if result := generateRandomElements(0); result != nil {
			t.Errorf("при размере 0 ожидался nil, получен %v", result)
		}
	})

	t.Run("нулевой размер даёт пустой слайс", func(t *testing.T) {
		if result := generateRandomElements(0); len(result) != 0 {
			t.Errorf("при размере 0 ожидалась длина 0, получена %d", len(result))
		}
	})

	t.Run("отрицательный размер даёт nil", func(t *testing.T) {
		if result := generateRandomElements(-10); result != nil {
			t.Errorf("при отрицательном размере ожидался nil, получен %v", result)
		}
	})

	t.Run("длина слайса равна запрошенному размеру", func(t *testing.T) {
		if result := generateRandomElements(100); len(result) != 100 {
			t.Errorf("ожидалась длина 100, получена %d", len(result))
		}
	})

	t.Run("значения лежат в диапазоне 1..1000000", func(t *testing.T) {
		result := generateRandomElements(1000)
		for i, v := range result {
			if v < 1 || v > 1_000_000 {
				t.Errorf("значение %d на индексе %d вне диапазона", v, i)
			}
		}
	})
}

func TestMaximum(t *testing.T) {
	t.Run("пустой слайс даёт -1", func(t *testing.T) {
		if got := maximum([]int{}); got != -1 {
			t.Errorf("ожидалось -1, получено %d", got)
		}
	})

	t.Run("nil слайс даёт -1", func(t *testing.T) {
		if got := maximum(nil); got != -1 {
			t.Errorf("ожидалось -1, получено %d", got)
		}
	})

	t.Run("один элемент — он и есть максимум", func(t *testing.T) {
		if got := maximum([]int{42}); got != 42 {
			t.Errorf("ожидалось 42, получено %d", got)
		}
	})

	t.Run("максимум в середине слайса", func(t *testing.T) {
		if got := maximum([]int{1, 5, 3, 9, 2}); got != 9 {
			t.Errorf("ожидалось 9, получено %d", got)
		}
	})

	t.Run("максимум в начале слайса", func(t *testing.T) {
		if got := maximum([]int{10, 3, 5, 2, 8}); got != 10 {
			t.Errorf("ожидалось 10, получено %d", got)
		}
	})

	t.Run("максимум в конце слайса", func(t *testing.T) {
		if got := maximum([]int{3, 5, 2, 8, 15}); got != 15 {
			t.Errorf("ожидалось 15, получено %d", got)
		}
	})

	t.Run("все элементы одинаковые", func(t *testing.T) {
		if got := maximum([]int{7, 7, 7, 7}); got != 7 {
			t.Errorf("ожидалось 7, получено %d", got)
		}
	})

	t.Run("отрицательные значения", func(t *testing.T) {
		if got := maximum([]int{-5, -2, -8, -3}); got != -2 {
			t.Errorf("ожидалось -2, получено %d", got)
		}
	})
}

func TestMaxChunks(t *testing.T) {
	t.Run("совпадает с maximum на длинах, не кратных 8", func(t *testing.T) {
		for n := 0; n <= 64; n++ {
			data := generateRandomElements(n)
			if n > 0 {
				data[n-1] = 1_000_001 // максимум заведомо в хвосте
			}
			want := maximum(data)
			if got := maxChunks(data); got != want {
				t.Errorf("для длины %d: maxChunks вернул %d, maximum — %d", n, got, want)
			}
		}
	})

	t.Run("хвост слайса не теряется", func(t *testing.T) {
		cases := []struct {
			data []int
			want int
		}{
			{nil, -1},
			{[]int{}, -1},
			{[]int{42}, 42},
			{[]int{1, 2, 3}, 3},
			{[]int{1, 2, 3, 4, 5, 6, 7}, 7},
			{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 8},
			{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9},
			{[]int{9, 1, 2, 3, 4, 5, 6, 7, 8}, 9},
			{[]int{5, 5, 5, 5, 5}, 5},
		}
		for _, c := range cases {
			if got := maxChunks(c.data); got != c.want {
				t.Errorf("maxChunks(%v) вернул %d, ожидалось %d", c.data, got, c.want)
			}
		}
	})

	t.Run("многократный вызов без гонок", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			data := generateRandomElements(100 + i%7)
			want := maximum(data)
			if got := maxChunks(data); got != want {
				t.Errorf("прогон %d: maxChunks вернул %d, maximum — %d", i, got, want)
			}
		}
	})
}
