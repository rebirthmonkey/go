package main

import (
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

func main() {
	// 创建 CVM 客户端
	credential := common.NewCredential("your-secret-id", "your-secret-key")
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"
	client, _ := cvm.NewClient(credential, "ap-guangzhou", cpf)

	// 创建 API 请求
	request := cvm.NewDescribeRegionsRequest()

	// 发送 API 请求并获取响应
	response, err := client.DescribeRegions(request)
	if err != nil {
		fmt.Println("failed to describe regions:", err)
		return
	}

	// 处理响应数据
	if len(response.Response.RegionSet) == 0 {
		fmt.Println("no regions found")
		return
	}
	fmt.Println("region list:")
	for _, region := range response.Response.RegionSet {
		fmt.Printf("- %s (%s)\n", *region.RegionName, *region.Region)
	}
}
