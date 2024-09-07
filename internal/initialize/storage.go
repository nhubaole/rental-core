package initialize

import (
	"context"
	"smart-rental/global"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func InitS3() {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}
	global.S3 = s3.NewFromConfig(sdkConfig)
}