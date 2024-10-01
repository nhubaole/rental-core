package services

import (
	"context"
	"errors"
	"fmt"
	"io"
	"smart-rental/global"
	"strings"

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
func (s *StorageSerivceImpl) UploadFile(bucketName string, objectKey string, data io.Reader, contentType string) (string, error) {
	// file, err := os.Open(fileName)
	// if err != nil {
	// 	fmt.Printf("Couldn't open file %v to upload. Here's why: %v\n", fileName, err)
	// } else {
	//defer file.Close()
	_, err := s.storage.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(objectKey),
		Body:        data,
		ContentType: &contentType,
	})
	if err != nil {
		fmt.Printf("Couldn't upload file to %v:%v. Here's why: %v\n",
			bucketName, objectKey, err)
	}
	//}

	presignClient := s3.NewPresignClient(s.storage)
	presignUrl, err := presignClient.PresignGetObject(context.Background(),
		&s3.GetObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(objectKey)},
	)
	if err != nil {
		// If there's an error getting the presigned URL, return the error.
		return "", err
	}
	urlList := strings.Split(presignUrl.URL, "?")
	return urlList[0], nil
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

// DeleteObject implements StorageSerivce.
func (s *StorageSerivceImpl) DeleteObject(bucketName string, objectKey string) error {
	_, err := s.storage.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
        Bucket: aws.String(bucketName),
        Key:    aws.String(objectKey),
    })
    if err != nil {
        return fmt.Errorf("unable to delete file %q from bucket %q, %v", objectKey, bucketName, err)
    }

    return nil
}