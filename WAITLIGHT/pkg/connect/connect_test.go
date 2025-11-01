package connect // 换成实际包名

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestGet(t *testing.T) {
	convey.Convey("基础用例", t, func() {
		url := "https://www.liwenzhou.com/posts/Go/unit-test-5/"
		got := Get(url)
		convey.So(got, convey.ShouldEqual, true)
	})

	convey.Convey("url请求不通的示例", t, func() {
		url := "posts/Go/unit-test-5/" // 故意给残缺路径
		got := Get(url)
		convey.So(got, convey.ShouldBeFalse)
	})
}