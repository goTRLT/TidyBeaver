package sources

import (
	"encoding/json"
	"log"
	"os/exec"
	models "tidybeaver/pkg/models"
)

func GetOSLogs() (osl models.OSLogs, err error) {
	output1, output2, output3, err := RunCommands()

	if err != nil {
		log.Fatal(err)
	}

	out, err := MergeOutput(output1, output2, output3)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(out, &osl.OSL)

	if err != nil {
		log.Fatal(err)
	}

	return osl, err
}

func RunCommands() (outputApp []byte, outputSys []byte, outputSec []byte, err error) {
	cmdApp := exec.Command("powershell", "-Command", "Get-EventLog -LogName Application -Newest 5 | ConvertTo-Json -Depth 2; ")
	outputApp, errApp := cmdApp.Output()

	if errApp != nil {
		log.Fatal("Error running PowerShell command: ", errApp)
		return outputApp, outputSys, outputSec, errApp
	}
	cmdSys := exec.Command("powershell", "-Command", "Get-EventLog -LogName System -Newest 5 | ConvertTo-Json -Depth 2; ")
	outputSys, errSys := cmdSys.Output()

	if errSys != nil {
		log.Fatal("Error running PowerShell command: ", errSys)
		return outputApp, outputSys, outputSec, errSys
	}

	cmdSec := exec.Command("powershell", "-Command", "Get-EventLog -LogName Security -Newest 5 | ConvertTo-Json -Depth 2")
	outputSec, errSec := cmdSec.Output()

	if errSec != nil {
		log.Fatal("Error running PowerShell command: ", errSec)
		return outputApp, outputSys, outputSec, errSec
	}
	return outputApp, outputSys, outputSec, err
}

func MergeOutput(outputs ...[]byte) ([]byte, error) {
	var out []any
	for _, vals := range outputs {
		var temp []any
		err := json.Unmarshal(vals, &temp)

		if err != nil {
			log.Fatal("Error unmarshaling JSON: ", err)
			return vals, err
		}
		out = append(out, temp...)
	}
	return json.Marshal(out)
}
