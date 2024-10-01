package services

import "io"

type StorageSerivce interface {
	UploadFile(bucketName string, objectKey string, data io.Reader, contentType string) (string, error)
	CreateBucket(bucketName string) error
	IsBucketExists(bucketName string) (bool, error)
	DeleteObject(bucketName string, objectKey string) error
}