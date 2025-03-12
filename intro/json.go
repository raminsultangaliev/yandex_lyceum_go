package main

import (
	"encoding/json"
)

type Student struct {
	Name string `json:"name"`
	Grade int `json:"grade"`
}

func modifyJSON(jsonData []byte) ([]byte, error) {
	var students []Student
	err := json.Unmarshal(jsonData, &students)
	if err != nil {
		return nil, err
	}

	for i := range students {
		students[i].Grade += 1
	}

	outputJSON, err := json.Marshal(students)
	if err != nil {
		return nil, err
	}
	return outputJSON, nil
}

func mergeJSONData(jsonDataList ...[]byte) ([]byte, error) {
	var mergedData []interface{}

	for _, jsonData := range jsonDataList {
		var data []interface{}
		err := json.Unmarshal(jsonData, &data)
		if err != nil {
			return nil, err
		}

		mergedData = append(mergedData, data...)
	}

	mergedJSON, err := json.Marshal(mergedData)
	if err != nil {
		return nil, err
	}

	return mergedJSON, nil
}

func splitJSONByClass(jsonData []byte) (map[string][]byte, error) {
	var data []map[string]interface{}
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}

	classMap := make(map[string][]interface{})

	for _, item := range data {
		class, _ := item["class"].(string)
		classMap[class] = append(classMap[class], item)
	}

	classJSON := make(map[string][]byte)

	for class, classData := range classMap {
		classJSON[class], err = json.Marshal(classData)
		if err != nil {
			return nil, err
		}
	}
	return classJSON, nil
}