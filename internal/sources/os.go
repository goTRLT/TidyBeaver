package sources

import (
	"encoding/json"
	"log"
	"os/exec"
	types "tidybeaver/pkg/types"
)

func FetchOSLogs() (OSLogs types.OSLogs, err error) {
	output1, output2, output3, err := RunCommands()

	if err != nil {
		log.Fatal(err)
	}

	out, err := MergeJSONOutputs(output1, output2, output3)
	// fmt.Println("out", string(out))
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(out, &OSLogs.OS)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Print("wel ", OSLogs)

	return OSLogs, err
}

func RunCommands() (outputApp []byte, outputSys []byte, outputSec []byte, err error) {
	// test := config.ConfigValues.App.LogAmount
	cmdApp := exec.Command("powershell", "-Command", "Get-EventLog -LogName Application -Newest 20 | ConvertTo-Json -Depth 2; ")
	outputApp, errApp := cmdApp.Output()

	if errApp != nil {
		log.Fatal("Error running PowerShell command: ", errApp)
		return outputApp, outputSys, outputSec, errApp
	}

	//fmt.Print("outputApp", string(outputApp))

	cmdSys := exec.Command("powershell", "-Command", "Get-EventLog -LogName System -Newest 20 | ConvertTo-Json -Depth 2; ")
	outputSys, errSys := cmdSys.Output()

	if errSys != nil {
		log.Fatal("Error running PowerShell command: ", errSys)
		return outputApp, outputSys, outputSec, errSys
	}

	//fmt.Print("outputSys", string(outputSys))

	cmdSec := exec.Command("powershell", "-Command", "Get-EventLog -LogName Security -Newest 20 | ConvertTo-Json -Depth 2")
	outputSec, errSec := cmdSec.Output()

	if errSec != nil {
		log.Fatal("Error running PowerShell command: ", errSec)
		return outputApp, outputSys, outputSec, errSec
	}

	//fmt.Print("outputSec", string(outputSec))

	return outputApp, outputSys, outputSec, err
}

func MergeJSONOutputs(outputs ...[]byte) ([]byte, error) {
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
