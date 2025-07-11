package models

import "time"

type FSLogs struct {
	FSLog []FSLog
}

type FSLog struct {
	Category           string    `json:"category,omitempty"`
	CategoryNumber     int       `json:"category_number,omitempty"`
	Checksum           string    `json:"checksum,omitempty"`
	ClientIP           string    `json:"client_ip,omitempty"`
	Column             string    `json:"column,omitempty"`
	Component          string    `json:"component,omitempty"`
	ComputerName       string    `json:"computer_name,omitempty"`
	Constraint         string    `json:"constraint,omitempty"`
	Container          string    `json:"container,omitempty"`
	CorrelationID      string    `json:"correlation_id,omitempty"`
	Data               []int     `json:"Data,omitempty"`
	Datatype           string    `json:"datatype,omitempty"`
	Detail             string    `json:"detail,omitempty"`
	Endpoint           string    `json:"endpoint,omitempty"`
	EntryType          int       `json:"entry_type,omitempty"`
	Environment        string    `json:"environment,omitempty"`
	Errcode            string    `json:"errcode,omitempty"`
	EventID            int       `json:"event_id,omitempty"`
	EventType          string    `json:"event_type,omitempty"`
	Path               string    `json:"file_path,omitempty"`
	FileSize           int64     `json:"file_size,omitempty"`
	Host               string    `json:"host,omitempty"`
	HTTPMethod         string    `json:"http_method,omitempty"`
	Index              int       `json:"index,omitempty"`
	InstanceID         int64     `json:"instance_id,omitempty"`
	LatencyMs          int       `json:"latency_ms,omitempty"`
	Level              string    `json:"level,omitempty"`
	LineNumber         int       `json:"line_number,omitempty"`
	LogName            string    `json:"log_name,omitempty"`
	MachineName        string    `json:"MachineName,omitempty"`
	Message            string    `json:"message,omitempty"`
	RequestBody        string    `json:"request_body,omitempty"`
	ReplacementStrings []string  `json:"replacement_strings,omitempty"`
	ResponseBody       string    `json:"response_body,omitempty"`
	RowsAffected       int       `json:"rows_affected,omitempty"`
	Schema             string    `json:"schema,omitempty"`
	Service            string    `json:"service,omitempty"`
	Source             string    `json:"source,omitempty"`
	SplitLines         string    `json:"SplitLines,omitempty"`
	SpanID             string    `json:"span_id,omitempty"`
	StatusCode         int       `json:"status_code,omitempty"`
	TableName          string    `json:"table_name,omitempty"`
	TimeGenerated      time.Time `json:"time_generated,omitempty"`
	TimeWritten        time.Time `json:"time_written,omitempty"`
	TransactionID      string    `json:"transaction_id,omitempty"`
	UserAgent          string    `json:"user_agent,omitempty"`
	UserID             string    `json:"user_id,omitempty"`
	UserName           string    `json:"user_name,omitempty"`
	Query              string    `json:"query,omitempty"`
}

func (v FSLog) ToAggregatedLog() AggregatedLog {
	return AggregatedLog{
		Category:           v.Category,
		CategoryNumber:     v.CategoryNumber,
		Checksum:           v.Checksum,
		ClientIP:           v.ClientIP,
		Column:             v.Column,
		Component:          v.Component,
		ComputerName:       v.ComputerName,
		Constraint:         v.Constraint,
		Container:          v.Container,
		CorrelationID:      v.CorrelationID,
		Data:               v.Data,
		Datatype:           v.Datatype,
		Detail:             v.Detail,
		Endpoint:           v.Endpoint,
		EntryType:          v.EntryType,
		Environment:        v.Environment,
		Errcode:            v.Errcode,
		EventID:            v.EventID,
		EventType:          v.EventType,
		Path:               v.Path,
		FileSize:           v.FileSize,
		Host:               v.Host,
		HTTPMethod:         v.HTTPMethod,
		Index:              v.Index,
		InstanceID:         v.InstanceID,
		LatencyMs:          v.LatencyMs,
		Level:              v.Level,
		LineNumber:         v.LineNumber,
		LogName:            v.LogName,
		MachineName:        v.MachineName,
		Message:            v.Message,
		RequestBody:        v.RequestBody,
		ReplacementStrings: v.ReplacementStrings,
		ResponseBody:       v.ResponseBody,
		RowsAffected:       v.RowsAffected,
		Schema:             v.Schema,
		Service:            v.Service,
		Source:             "FileSystem: " + v.Source,
		SplitLines:         v.SplitLines,
		SpanID:             v.SpanID,
		StatusCode:         v.StatusCode,
		TableName:          v.TableName,
		TimeGenerated:      v.TimeGenerated,
		TimeWritten:        time.Now(),
		TransactionID:      v.TransactionID,
		UserAgent:          v.UserAgent,
		UserID:             v.UserID,
		UserName:           v.UserName,
		Query:              v.Query,
	}
}
