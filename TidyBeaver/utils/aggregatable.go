package utils

import "tidybeaver/pkg/models"

type Aggregatable interface {
	ToAggregatedLog() models.AggregatedLog
}