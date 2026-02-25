package urltool_test

import (
	"shortener/pkg/urltool"
	"testing"
)

func TestGetBasePath(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		targetUrl string
		want      string
		wantErr   bool
	}{
		// TODO: Add test cases.
		{name: "valid url with path", targetUrl: "http://example.com/path/to/resource", want: "resource", wantErr: false},
		{name: "url with missing host", targetUrl: "http:///path/to/resource", want: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := urltool.GetBasePath(tt.targetUrl)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetBasePath() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetBasePath() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("GetBasePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
