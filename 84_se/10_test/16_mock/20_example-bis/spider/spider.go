package spider

//go:generate mockgen -destination mock/mock_spider.go -package spider github.com/rebirthmonkey/go/84_software-engineering/10_test/10_unit/30_mock/20_example/spider Spider

type Spider interface {
	GetBody() string
}
