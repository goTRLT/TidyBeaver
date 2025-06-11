package utils

import "tidybeaver/pkg/models"

func TransformSlice[T Aggregatable](logs []T) []models.AggregatedLog {
    var result []models.AggregatedLog
    for _, log := range logs {
        result = append(result, log.ToAggregatedLog())
    }
    return result
}