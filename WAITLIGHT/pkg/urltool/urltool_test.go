package urltool_test

import(
	"testing"
	"ztt1/pkg/urltool"
)

func TestGetBaseS(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		targeturl string
		want      string
		wantErr   bool
	}{
		// TODO: Add test cases.
		{name: "基本示例", targeturl: "https://www.example.com/abc/", want: "abc", wantErr: false},
		{name: "无效示例", targeturl: "/wwww/abc/", want: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := urltool.GetBasePath(tt.targeturl)
			t.Logf("got len=%d bytes: %#v", len(got), []byte(got))
			t.Logf("want len=%d bytes: %#v", len(tt.want), []byte(tt.want))
			t.Logf("got=|%q|, want=|%q|", got, tt.want) // %q 会把不可见字符转义
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
			if got!=tt.want{
				t.Errorf("GetBasePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
