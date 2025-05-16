package aggregator

import (
	"tidybeaver/pkg/models"
)

func TransformSampleLogs(MockLogs *models.SampleLogs)      {}
func TransformAPILogs(APILogs *[]string)                   {}
func TransformDBLogs(DBLogs *[]string)                     {}
func TransformFSLogs(TransformedLogs *models.StandardLogs) {}
func TransformMSVLogs(MSVLogs *[]string)                   {}
func TransformOSLogs(OSLogs *models.OSLogs)                {}
