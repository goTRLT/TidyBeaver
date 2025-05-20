package aggregator

import (
	"log"
	config "tidybeaver/internal/config"
	source "tidybeaver/internal/sources"
	storage "tidybeaver/internal/storage"
	"tidybeaver/pkg/models"
)

var AggregatedLogs models.AggregatedLogs
var SampleLogs models.SampleLogs
var OSLogs models.OSLogs
var FSLogs models.FSLogs
var APILogs []string //Placeholder
var MSVLogs []string //Placeholder
var DBLogs models.DBLogs

func Init() {
	FetchLogs()
	TransformLogs()
	SaveLogs()
}

// Refactor to Helper Function
func FetchLogs() {
	var err error
	if config.UserInputConfigValues.UseSampleLogs {
		SampleLogs, err = source.CreateSampleLogs()

		if err != nil {
			log.Fatal(err)
		}

	} else {
		if config.UserInputConfigValues.UseAPI {
			//TODO
		}
		if config.UserInputConfigValues.UseDatabase {
			DBLogs, err = source.FetchDBLogs()

			if err != nil {
				log.Fatal(err)
			}

		}
		if config.UserInputConfigValues.UseFileSystem {
			FSLogs, err = source.FetchFSLogs()

			if err != nil {
				log.Fatal(err)
			}

		}
		if config.UserInputConfigValues.UseMicroservice {
			//TODO
		}
		if config.UserInputConfigValues.UseWindowsEvents {

			OSLogs, err = source.FetchOSLogs()

			if err != nil {
				log.Fatal(err)
			}

		}
	}
}

func TransformLogs() {
	if len(SampleLogs.SampleLog) != 0 {
		TransformedLogs := TransformSampleLogs(&SampleLogs)
		Aggregate(&TransformedLogs)
	}
	if len(OSLogs.OS) != 0 {
		TransformedLogs := TransformOSLogs(&OSLogs)
		Aggregate(&TransformedLogs)
	}
	if len(FSLogs.FSLog) != 0 {
		TransformedLogs := TransformFSLogs(&FSLogs)
		Aggregate(&TransformedLogs)
	}
	if len(DBLogs.DBLog) != 0 {
		TransformedLogs := TransformDBLogs(&DBLogs)
		Aggregate(&TransformedLogs)
	}
	// if len(APILogs) != 0 {
	// 	//TODO
	// }
	// if len(MSVLogs) != 0 {
	// 	//TODO
	// }
}

func SaveLogs() {
	storage.SaveLogsJson(&AggregatedLogs)
	storage.DBInsertLogs(&AggregatedLogs)
}

func Aggregate(transformedLog *[]models.AggregatedLog) {
	AggregatedLogs.AggregatedLogSlice = append(AggregatedLogs.AggregatedLogSlice, *transformedLog...)
	// fmt.Println(AggregatedLogs.AggregatedLogSlice)
}
