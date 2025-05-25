package types

type OSLogs struct {
	OS []OS
}

type OS struct {
	Category           string   `json:"Category"`
	CategoryNumber     int      `json:"CategoryNumber"`
	Container          string   `json:"Container"`
	Data               []int    `json:"Data"`
	EntryType          int      `json:"EntryType"`
	EventID            int      `json:"EventID"`
	Index              int      `json:"Index"`
	InstanceID         int64    `json:"InstanceID"`
	MachineName        string   `json:"MachineName"`
	Message            string   `json:"Message"`
	ReplacementStrings []string `json:"ReplacementStrings"`
	Source             string   `json:"Source"`
	SplitLines         string   `json:"SplitLines"`
	TimeGenerated      string   `json:"TimeGenerated"`
	TimeWritten        string   `json:"TimeWritten"`
	UserName           string   `json:"UserName"`
}
