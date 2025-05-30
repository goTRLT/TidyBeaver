package s3

import (
	"fmt"
	config "tidybeaver/internal/config"
)

func InitS3() {
	buckets := ListBuckets()
	if buckets.Buckets == nil {
		CreateBucket()
		buckets = ListBuckets()
	}

	if buckets.Buckets != nil {
		fmt.Println("Do you wish to download the Previous Logs already stored on your s3 bucket before uploading the New Logs?")
		if config.CheckAnswer() {
			DownloadLogs()
		}
	}
	UploadLogs()
}
