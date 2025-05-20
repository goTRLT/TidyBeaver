package sources

import (
	"encoding/json"
	"net/http"
	"tidybeaver/internal/config"
	"tidybeaver/pkg/models"
)

func FetchAPILogs() (APILogs []models.APILogs, err error) {

	//TODO Add Timeout
	resp, err := http.Get(config.ConfigValues.API.BaseURL + config.ConfigValues.App.LogAmount)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// fmt.Println("Resp ", resp)

	var responses []models.APILog

	if err := json.NewDecoder(resp.Body).Decode(&responses); err != nil {
		return nil, err
	}

	// fmt.Println("ActResp ", responses)

	append(APILogs, responses...)

	return APILogs, err
}
