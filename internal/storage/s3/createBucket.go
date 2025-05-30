package s3

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func CreateBucket() {
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
	s3Client := s3.New(newSession)
	bucketName := "tidybeaverbucket"

	err = Create(s3Client, bucketName)

	if err != nil {
		log.Fatal(err)
		return
	}
}

func Create(client *s3.S3, bucketName string) error {
	_, err := client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})
	return err
}
