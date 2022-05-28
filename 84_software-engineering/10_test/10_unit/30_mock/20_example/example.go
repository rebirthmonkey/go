package gomock

import (
	"github.com/rebirthmonkey/go/84_software-engineering/10_test/10_unit/30_mock/20_example/spider"
)

func GetGoVersion(s spider.Spider) string {
	body := s.GetBody()
	return body
}
