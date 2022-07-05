package service

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	image "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/image/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/image/v2/model"
	region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/image/v2/region"
)

func s() {
	ak := "5ZKDXQSJ2LEBVWGY6ADM"
	sk := "XSJPKUnWYJYsGO3je1LgdeINHrQiwOQ784B2UyrG"

	auth := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		Build()

	client := image.NewImageClient(
		image.ImageClientBuilder().
			WithRegion(region.ValueOf("cn-east-3")).
			WithCredential(auth).
			Build())

	request := &model.RunImageTaggingRequest{}
	limitImageTaggingReq := int32(50)
	thresholdImageTaggingReq := float32(95)
	languageImageTaggingReq := "zh"
	urlImageTaggingReq := "https://ts1.cn.mm.bing.net/th/id/R-C.5d4c1595c19c670dab838e322ecb681a?rik=1uMbbERSYuEGsw&riu=http%3a%2f%2fwww.dnzhuti.com%2fuploads%2fallimg%2f180313%2f95-1P313154R5.jpg&ehk=LjNKD8D%2bYzuuANUZXHiLeT%2bACwqjCjOgopY6tMbf35Q%3d&risl=&pid=ImgRaw&r=0"
	request.Body = &model.ImageTaggingReq{
		Limit:     &limitImageTaggingReq,
		Threshold: &thresholdImageTaggingReq,
		Language:  &languageImageTaggingReq,
		Url:       &urlImageTaggingReq,
	}
	response, err := client.RunImageTagging(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
	}
}
