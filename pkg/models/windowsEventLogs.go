package models

type WindowsEventLogs struct {
	WindowsEventLogs []WindowsEventLog
}

type WindowsEventLog struct {
	MachineName        string   `json:"MachineName"`
	Data               []string `json:"Data"`
	Index              int      `json:"Index"`
	Category           string   `json:"Category"`
	CategoryNumber     int      `json:"CategoryNumber"`
	EventID            int      `json:"EventID"`
	EntryType          int      `json:"EntryType"`
	Message            string   `json:"Message"`
	Source             string   `json:"Source"`
	ReplacementStrings []string `json:"ReplacementStrings"`
	InstanceID         int      `json:"InstanceID"`
	TimeGenerated      string   `json:"TimeGenerated"`
	TimeWritten        string   `json:"TimeWritten"`
	UserName           string   `json:"UserName"`
	SplitLines         string   `json:"SplitLines"`
	Container          string   `json:"Container"`
}
