package models

import "time"

type DBLogs struct {
	DBLog []DBLog
}

type DBLog struct {
	Level      string `json:"level"`
	Column     string `json:"column"`
	Constraint string `json:"constraint"`
	Datatype   string `json:"datatype"`
	Table_name string `json:"table_name"`
	Schema     string `json:"schema"`
	Errcode    string `json:"errcode"`
	Detail     string `json:"detail"`
}

func (v DBLog) ToAggregatedLog() AggregatedLog {
	return AggregatedLog{
		Column:        v.Column,
		CorrelationID: v.Constraint,
		Datatype:      v.Datatype,
		Detail:        v.Detail,
		Errcode:       v.Errcode,
		Level:         v.Level,
		Schema:        v.Schema,
		Source:        "Database",
		TableName:     v.Table_name,
		TimeGenerated: time.Now(),
		TimeWritten:   time.Now(),
	}
}
