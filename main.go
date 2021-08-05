package main

import (
	"cis-image-sync/pkg/transfer"
	"fmt"
	"tkestack.io/image-transfer/pkg/utils"
)

func main() {
	// TODO consumer msg from mq
	job, err := GenerateTransferJob("cis-hub-huabei-3.cmecloud.cn/image_test/zhouyu:df", "cis-hub-huadong-4.cmecloud.cn/test6/aaa:v1.0")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if err := job.Run(); err != nil {
		fmt.Println(err.Error())
	}
}


// GenerateTransferJob creates transfer jobs from normalURLPair
func GenerateTransferJob(source string, target string) (*transfer.Job, error) {
	if source == "" {
		return nil, fmt.Errorf("source url should not be empty")
	}

	sourceURL, err := utils.NewRepoURL(source)
	if err != nil {
		return nil, fmt.Errorf("url %s format error: %v", source, err)
	}

	if target == "" {
		return nil, fmt.Errorf("target url should not be empty")
	}

	targetURL, err := utils.NewRepoURL(target)
	if err != nil {
		return nil, fmt.Errorf("url %s format error: %v", target, err)
	}

	// if tag is not specific
	if sourceURL.GetTag() == "" {
		return nil, fmt.Errorf("source tag empty, source: %s", sourceURL.GetURL())
	}

	if targetURL.GetTag() == "" {
		return nil, fmt.Errorf("target tag empty, target: %s", targetURL.GetURL())
	}

	var imageSource *transfer.ImageSource
	var imageTarget *transfer.ImageTarget

	imageSource, err = transfer.NewImageSource(sourceURL.GetRegistry(), sourceURL.GetRepoWithNamespace(), sourceURL.GetTag(), "zhounanjun", "xxxx", true)
	if err != nil {
		return nil, fmt.Errorf("generate %s image source error: %v", sourceURL.GetURL(), err)
	}

	imageTarget, err = transfer.NewImageTarget(targetURL.GetRegistry(), targetURL.GetRepoWithNamespace(), targetURL.GetTag(), "zhounanjun", "xxxx", true)
	if err != nil {
		return nil, fmt.Errorf("generate %s image target error: %v", sourceURL.GetURL(), err)
	}


	return transfer.NewJob(imageSource, imageTarget), nil
}