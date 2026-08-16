package io

import "os"


func SaveToFile(data []byte, filename string) error {

	return os.WriteFile(filename, data, 0644)
}

func ReadFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}
	
	return data, nil
}
