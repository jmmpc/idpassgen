package idpassgen

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var n = 32
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
var runeset = []rune(chars)

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

func BenchmarkNewString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewString(n, runeset, rnd)
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

func TestNewString(t *testing.T) {
	tt := []struct {
		inputLength  int
		outputLength int
		charset      []rune
	}{
		{0, 1, []rune(chars[33:59])},
		{-13, 1, runeset},
		{32, 32, runeset},
		{3, 3, runeset},
		{1, 1, runeset},
		{12, 12, []rune("0123456789абвгґдеєжзиіїйклмнопрстуфхцчшщьюя")},
	}

	for _, tc := range tt {
		result := NewString(tc.inputLength, tc.charset, rnd)
		fmt.Printf("Generated string: %s\n", result)
		if l := len([]rune(result)); l != tc.outputLength {
			t.Errorf("expected result lendth %d; got %d", tc.outputLength, l)
		}
	}
}
