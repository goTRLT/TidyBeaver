package s3

func InitS3() {
	
	buckets := ListBuckets()
	if buckets.Buckets == nil {
		CreateBucket()
		buckets = ListBuckets()
	}

	if buckets.Buckets != nil {
		DownloadLogs()
	}
	UploadLogs()
}
