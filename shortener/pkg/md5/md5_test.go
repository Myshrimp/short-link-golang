package md5_test

import(
	"shortener/pkg/md5"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		data []byte
		want string
	}{
		// TODO: Add test cases.
		{name: "empty data", data: []byte(""), want: "d41d8cd98f00b204e9800998ecf8427e"},
		{name: "hello world", data: []byte("hello world"), want: "5eb63bbbe01eeed093cb22bb8f5acdc3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := md5.Sum(tt.data)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
