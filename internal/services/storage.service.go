package services

type StorageSerivce interface {
	UploadFile(bucketName string, objectKey string, fileName string) error
	CreateBucket(bucketName string) error
	IsBucketExists(bucketName string) (bool, error)
	
}