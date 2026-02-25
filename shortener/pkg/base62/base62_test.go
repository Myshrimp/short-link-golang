package base62_test

import (
	"shortener/pkg/base62"
	"testing"
)

func TestInt2String(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		num  uint64
		want string
	}{
		// TODO: Add test cases.
		{name: "zero", num: 0, want: "0"},
		{name: "one", num: 1, want: "1"},
		{name: "ten", num: 10, want: "a"},
		{name: "sixty-one", num: 61, want: "Z"},
		{name: "sixty-two", num: 62, want: "10"},
		{name: "one hundred twenty-three", num:6347, want: "1En"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := base62.Int2String(tt.num)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("Int2String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString2Int(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		str  string
		want uint64
	}{
		// TODO: Add test cases.
		{name: "zero", str: "0", want: 0},
		{name: "one", str: "1", want: 1},
		{name: "ten", str: "a", want: 10},
		{name: "sixty-one", str: "Z", want: 61},
		{name: "sixty-two", str: "10", want: 62},
		{name: "one hundred twenty-three", str: "1En", want: 6347},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := base62.String2Int(tt.str)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("String2Int() = %v, want %v", got, tt.want)
			}
		})
	}
}
