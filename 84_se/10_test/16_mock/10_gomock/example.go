package gomock

import (
	"github.com/rebirthmonkey/go/84_se/10_test/16_mock/10_gomock/spider"
)

func GetGoVersion(s spider.Spider) string {
	body := s.GetBody()
	return body
}
