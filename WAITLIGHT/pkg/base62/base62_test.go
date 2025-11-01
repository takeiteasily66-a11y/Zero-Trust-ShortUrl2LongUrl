package base62_test

import(
	"testing"
	"ztt1/pkg/base62"
)

func TestString2Int(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want uint64
	}{
		// TODO: Add test cases.
		{name: "0",s:"0",want: 0},
		{name: "10",s:"10",want: 62},
		{name: "1En",s:"1En",want: 6347},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := base62.String2Int(tt.s)
			// TODO: update the condition below to compare got with tt.want.
			if got!=tt.want{
				t.Errorf("String2Int() = %v, want %v", got, tt.want)
			}
		})
	}
}
