package s3

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func DownloadLogs() {
	newSession, err := session.NewSessionWithOptions(session.Options{
		Profile: "TRLTech",
		Config: aws.Config{
			Region: aws.String("us-west-1"),
		},
	})

	if err != nil {
		log.Println(err)
		return
	}

	downloader := s3manager.NewDownloader(newSession)
	bucketName := "tidybeaverbucket"

	key := os.Getenv("LOGS_FILE_NAME")
	err = Download(downloader, bucketName, key)

	if err != nil {
		log.Println(err)
		return
	}
}

func Download(downloader *s3manager.Downloader, bucketName string, key string) error {
	file, err := os.Create(key)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})
	return err
}
