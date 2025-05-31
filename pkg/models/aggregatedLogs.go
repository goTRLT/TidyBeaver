package models

import "time"

type AggregatedLogs struct {
	AL []AggregatedLog
}

type AggregatedLog struct {
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
