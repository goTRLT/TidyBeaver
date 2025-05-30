package s3

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func ListBuckets() {
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

	buckets, err := List(s3Client)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Buckets found in AWS: ")
	for _, bucket := range buckets.Buckets {
		fmt.Printf("Name: %s , Creation Date: %s\n", *bucket.Name, *bucket.CreationDate)
	}
}

func List(client *s3.S3) (*s3.ListBucketsOutput, error) {
	res, err := client.ListBuckets(nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}
