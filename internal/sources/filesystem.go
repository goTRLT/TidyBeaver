package sources

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func getFileNameWithoutExtension(filePath string) string {
	// Get the base name of the file (without the directory path)
	base := filepath.Base(filePath)

	// Split the base name into name and extension
	name := strings.TrimSuffix(base, filepath.Ext(base))

	return name
}

func getParentDirectory(filePath string) string {
	// Get the directory of the file
	dir := filepath.Dir(filePath)

	// Get the parent directory
	parentDir := filepath.Dir(dir)

	return parentDir
}

func parseJsonFile(filePath string) (map[string]interface{}, error) {
	// Open the JSON file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file content
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Parse the JSON data into a map
	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func syncFile(filePath string, data map[string]interface{}) error {
	// Open the file for writing
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Convert the data to JSON format
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// Write the JSON data to the file
	_, err = file.Write(jsonData)

	if err != nil {
		return err
	}

	return nil
}
