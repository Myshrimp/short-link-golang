package connect

import(
	"testing"
	"github.com/smartystreets/goconvey/convey"
)

func TestGet(t *testing.T) {
	convey.Convey("TestGet", t, func() {
		convey.Convey("valid url", func() {
			url := "http://www.baidu.com"
			result := Get(url)
			convey.So(result, convey.ShouldBeTrue)
		})
		convey.Convey("invalid url", func() {
			url := "http://www.invalidurl.com"
			result := Get(url)
			convey.So(result, convey.ShouldBeFalse)
		})
	})
}
