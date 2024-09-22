package initialize

import (
	"context"
	"smart-rental/global"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func InitS3() {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(global.Config.S3.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			global.Config.S3.AccessKeyID, 
			global.Config.S3.SecretKey, 
			"")),
	)

	if err != nil {
		panic(err)
	}
	global.S3 = s3.NewFromConfig(cfg)
}
