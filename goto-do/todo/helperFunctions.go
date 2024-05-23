package todo

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
)

func randomString() string {
	id := uuid.New()

	return string(id.String())
}

func isJSONFileEmpty(filePath string) (bool, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return false, err
	}

	// Check file size
	if info.Size() == 0 {
		return true, nil
	}

	// Read file content
	data, err := os.ReadFile(filePath)
	if err != nil {
		return false, err
	}

	// Check if it's an empty JSON object or array
	var obj interface{}
	err = json.Unmarshal(data, &obj)
	if err != nil {
		return false, err
	}

	switch obj.(type) {
	case map[string]interface{}:
		return len(obj.(map[string]interface{})) == 0, nil
	case []interface{}:
		return len(obj.([]interface{})) == 0, nil
	default:
		return false, fmt.Errorf("unexpected JSON type")
	}
}

func createJsonEmptyArrayFile(filePath string) {
	var emptyArray []interface{}

	jsonData, err := json.Marshal(emptyArray)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	err = os.WriteFile(filePath, jsonData, 0775)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
