package services

import (
	"context"
	"errors"
	"fmt"
	"os"
	"smart-rental/global"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
)

var region = global.Config.S3

type StorageSerivceImpl struct {
	storage *s3.Client
}


func NewStorageServiceImpl() StorageSerivce {
	return &StorageSerivceImpl{
		storage: global.S3,
	}
}

// UploadFile implements StorageSerivce.
func (s *StorageSerivceImpl) UploadFile(bucketName string, objectKey string, fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Couldn't open file %v to upload. Here's why: %v\n", fileName, err)
	} else {
		defer file.Close()
		_, err = s.storage.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(objectKey),
			Body:   file,
		})
		if err != nil {
			fmt.Printf("Couldn't upload file %v to %v:%v. Here's why: %v\n",
				fileName, bucketName, objectKey, err)
		}
	}
	return err
}

func (s *StorageSerivceImpl) CreateBucket(bucketName string) error {
	_, err := s.storage.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraint(region.Region),
		},
	})
	if err != nil {
		fmt.Printf("Couldn't create bucket %v in Region %v. Here's why: %v\n",
			bucketName, region, err)
	}
	return err
}

// IsBucketExists implements StorageSerivce.
func (s *StorageSerivceImpl) IsBucketExists(bucketName string) (bool, error) {
	_, err := s.storage.HeadBucket(context.TODO(), &s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})
	exists := true
	if err != nil {
		var apiError smithy.APIError
		if errors.As(err, &apiError) {
			switch apiError.(type) {
			case *types.NotFound:
				fmt.Printf("Bucket %v is available.\n", bucketName)
				exists = false
				err = nil
			default:
				fmt.Printf("Either you don't have access to bucket %v or another error occurred. "+
					"Here's what happened: %v\n", bucketName, err)
			}
		}
	} else {
		fmt.Printf("Bucket %v exists and you already own it.", bucketName)
	}

	return exists, err
}
