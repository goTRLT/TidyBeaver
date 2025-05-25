package storage

import (
	"encoding/json"
	"fmt"
	"os"
	types "tidybeaver/pkg/types"
)

func JSONSaveLogs(Logs *types.AggregatedLogs) {
	encodedLogs, err := json.Marshal(Logs)
	if err != nil {
		return
	} else {
		os.WriteFile((`.\Logs\TidyBeaverAdaptedLogs.json`), encodedLogs, 0644)
		fmt.Println("Logs saved as Json")
	}
}
