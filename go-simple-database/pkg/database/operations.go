package database

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/datarohit/go-database/pkg/utils"
)

func (driver *Driver) Write(collection string, resource string, data interface{}) error {
	if collection == "" {
		return fmt.Errorf("missing collection - no place to save record")
	}

	if resource == "" {
		return fmt.Errorf("missing resource - unable to save record (no name)")
	}

	mutex := driver.GetOrCreateMutex(collection)
	mutex.Lock()
	defer mutex.Unlock()

	dir := filepath.Join(driver.Dir, collection)
	finalPath := filepath.Join(dir, resource+".json")
	tempPath := finalPath + ".tmp"

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	jsonByteStr, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	jsonByteStr = append(jsonByteStr, byte('\n'))

	if err := os.WriteFile(tempPath, jsonByteStr, 0644); err != nil {
		return err
	}

	return os.Rename(tempPath, finalPath)
}

func (driver *Driver) Read(collection string, resource string, data interface{}) error {
	if collection == "" {
		return fmt.Errorf("missing collection - unable to read")
	}

	if resource == "" {
		return fmt.Errorf("missing resource - unable to rea record (no name)")
	}

	record := filepath.Join(driver.Dir, collection, resource)
	if _, err := utils.Stat(record); err != nil {
		return err
	}

	jsonByteStr, err := os.ReadFile(record + ".json")
	if err != nil {
		return err
	}

	return json.Unmarshal(jsonByteStr, &data)
}

func (driver *Driver) ReadAll(collection string) ([]string, error) {
	if collection == "" {
		return nil, fmt.Errorf("missing collection - unable to read")
	}

	dir := filepath.Join(driver.Dir, collection)
	if _, err := utils.Stat(dir); err != nil {
		return nil, err
	}

	files, _ := os.ReadDir(dir)

	var records []string
	for _, file := range files {
		jsonByteStr, err := os.ReadFile(filepath.Join(dir, file.Name()))
		if err != nil {
			return nil, err
		}

		records = append(records, string(jsonByteStr))
	}

	return records, nil
}

func (driver *Driver) Delete(collection string, resource string) error {
	path := filepath.Join(collection, resource)

	mutex := driver.GetOrCreateMutex(collection)
	mutex.Lock()
	defer mutex.Unlock()

	dir := filepath.Join(driver.Dir, path)

	switch fileInfo, err := utils.Stat(dir); {
	case fileInfo == nil, err != nil:
		return fmt.Errorf("unable to find file or directory named %v\n", path)

	case fileInfo.Mode().IsDir():
		return os.RemoveAll(dir)

	case fileInfo.Mode().IsRegular():
		return os.RemoveAll(dir + ".json")
	}
	return nil
}
