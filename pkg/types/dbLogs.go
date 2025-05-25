package types

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
