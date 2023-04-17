package spider

//go:generate mockgen -destination mock/mock_spider.go -package spider github.com/rebirthmonkey/go/84_se/10_test/16_mock/20_example/spider Spider

type Spider interface {
	GetBody() string
}
