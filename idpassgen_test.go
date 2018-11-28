package idpassgen

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var n = 32
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func BenchmarkNewID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewID(n, rnd)
	}
}

func BenchmarkNewPassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewPassword(n, rnd)
	}
}

func BenchmarkNewHex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewHex(n, rnd)
	}
}

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(n, chars, rnd)
	}
}

func TestNewID(t *testing.T) {
	tt := []struct {
		inputLength  int
		outputLength int
	}{
		{0, 1},
		{-13, 1},
		{32, 32},
		{3, 3},
		{1, 1},
	}

	for _, tc := range tt {
		result := NewID(tc.inputLength, rnd)
		if l := len(result); l != tc.outputLength {
			t.Errorf("expected result lendth %d; got %d", tc.outputLength, l)
		}
	}
}

func TestNewPassword(t *testing.T) {
	tt := []struct {
		inputLength  int
		outputLength int
	}{
		{0, 4},
		{-13, 4},
		{32, 32},
		{3, 4},
		{1, 4},
	}

	for _, tc := range tt {
		result := NewPassword(tc.inputLength, rnd)
		if l := len(result); l != tc.outputLength {
			t.Errorf("expected result lendth %d; got %d", tc.outputLength, l)
		}
	}
}

func TestNewHex(t *testing.T) {
	tt := []struct {
		inputLength  int
		outputLength int
	}{
		{0, 1},
		{-13, 1},
		{32, 32},
		{3, 3},
		{1, 1},
	}

	for _, tc := range tt {
		result := NewHex(tc.inputLength, rnd)
		if l := len(result); l != tc.outputLength {
			t.Errorf("expected result lendth %d; got %d", tc.outputLength, l)
		}
	}
}

func TestString(t *testing.T) {
	tt := []struct {
		inputLength  int
		outputLength int
		charset      string
	}{
		{0, 1, chars[33:59]},
		{-13, 1, chars},
		{32, 32, chars},
		{3, 3, chars},
		{1, 1, chars},
		{12, 12, "0123456789абвгґдеєжзиіїйклмнопрстуфхцчшщьюя"},
	}

	for _, tc := range tt {
		result := String(tc.inputLength, tc.charset, rnd)
		fmt.Printf("Generated string: %s\n", result)
		if l := len([]rune(result)); l != tc.outputLength {
			t.Errorf("expected result lendth %d; got %d", tc.outputLength, l)
		}
	}
}
