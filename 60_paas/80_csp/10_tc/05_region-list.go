package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rebirthmonkey/go/pkg/tencentcloud"
	"github.com/rebirthmonkey/go/pkg/tencentcloud/region"
)

func main() {
	cli, _ := region.NewClient(
		tencentcloud.SecretId(os.Getenv("TENCENTCLOUD_SECRET_ID")),
		tencentcloud.SecretKey(os.Getenv("TENCENTCLOUD_SECRET_KEY")),
	)

	regions, _ := cli.ListRegions()
	data, _ := json.MarshalIndent(regions, "", "\t")
	fmt.Println(string(data))
}
