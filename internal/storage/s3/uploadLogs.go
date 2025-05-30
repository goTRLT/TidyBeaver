package s3

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadLogs() {
	newSession, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region: aws.String("us-west-1"),
		},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	uploader := s3manager.NewUploader(newSession)
	bucketName := "tidybeaverbucket"

	filePath := os.Getenv("LOGS_FOLDER_PATH")
	fileName := os.Getenv("LOGS_FILE_NAME")
	err = Upload(uploader, filePath, bucketName, fileName)

	if err != nil {
		log.Fatal(err)
		return
	}
}

func Upload(uploader *s3manager.Uploader, filePath string, bucketName string, fileName string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   file,
	})
	return err
}
