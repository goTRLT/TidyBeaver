package sources

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	models "tidybeaver/pkg/models"
)

func FetchOSLogs() (windowsEventLogs models.OSLogs, err error) {

	output1, output2, output3, err := RunCommands()

	if err != nil {
		log.Fatal(err)
	}

	out, err := MergeJSONOutputs(output1, output2, output3)
	// fmt.Println("out", string(out))
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(out, &windowsEventLogs.OS)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Print("wel", windowsEventLogs)

	return windowsEventLogs, err
}

func RunCommands() (outputApp []byte, outputSys []byte, outputSec []byte, err error) {
	cmdApp := exec.Command("powershell", "-Command", "Get-EventLog -LogName Application -Newest 2 | ConvertTo-Json -Depth 2; ")
	outputApp, errApp := cmdApp.Output()

	if errApp != nil {
		fmt.Printf("Error running PowerShell command: %v\n", errApp)
		return outputApp, outputSys, outputSec, errApp
	}

	//fmt.Print("outputApp", string(outputApp))

	cmdSys := exec.Command("powershell", "-Command", "Get-EventLog -LogName System -Newest 2 | ConvertTo-Json -Depth 2; ")
	outputSys, errSys := cmdSys.Output()

	if errSys != nil {
		fmt.Printf("Error running PowerShell command: %v\n", errSys)
		return outputApp, outputSys, outputSec, errSys
	}

	//fmt.Print("outputSys", string(outputSys))

	cmdSec := exec.Command("powershell", "-Command", "Get-EventLog -LogName Security -Newest 2 | ConvertTo-Json -Depth 2")
	outputSec, errSec := cmdSec.Output()

	if errSec != nil {
		fmt.Printf("Error running PowerShell command: %v\n", errSec)
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
			fmt.Printf("Error unmarshaling JSON: %v\n", err)
			return vals, err
		}
		for _, v := range temp {
			out = append(out, v)
		}
	}
	return json.Marshal(out)
}
