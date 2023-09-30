package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"
	"perfe/nginxbuilder/src/types"
)

func LoadConfigsFromDir(directory string) ([]types.Config, error) {
	var configs []types.Config

	// List files/folders in the directory
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	// Loop through each file/folder in the directory
	for _, file := range files {
		filePath := filepath.Join(directory, file.Name())

		if file.IsDir() {
			// If it's a directory, read the index.json for meta-information
			configFilePath := filepath.Join(filePath, "index.json")
			config, err := readConfigFile(configFilePath)
			if err != nil {
				return nil, err
			}

			// Read the items inside the /items/ folder
			itemsDir := filepath.Join(filePath, "items")
			itemFiles, _ := ioutil.ReadDir(itemsDir)
			for _, itemFile := range itemFiles {
				itemFilePath := filepath.Join(itemsDir, itemFile.Name())
				items, err := readItemsFile(itemFilePath)
				if err != nil {
					return nil, err
				}
				config.Items = append(config.Items, items...)
			}
			configs = append(configs, config)

		} else {
			// If it's a file, directly read its content as a Config with embedded items
			config, err := readConfigFile(filePath)
			if err != nil {
				return nil, err
			}
			configs = append(configs, config)
		}
	}

	return configs, nil
}

func readConfigFile(filePath string) (types.Config, error) {
	var config types.Config
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, errors.New("error parsing JSON in " + filePath + ": " + err.Error())
	}

	return config, nil
}

func readItemsFile(filePath string) ([]types.Element, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var rawItems []map[string]interface{}
	err = json.Unmarshal(data, &rawItems)
	if err != nil {
		return nil, err
	}

	var items []types.Element
	for _, rawItem := range rawItems {
		// Check if it's a server based on presence of specific fields
		if _, exists := rawItem["name"]; exists {
			serverBytes, _ := json.Marshal(rawItem)
			var server types.Server
			json.Unmarshal(serverBytes, &server)
			server.Init()
			items = append(items, server)
		} else if _, exists := rawItem["docs"]; exists {
			commentBytes, _ := json.Marshal(rawItem)
			var comment types.Comment
			json.Unmarshal(commentBytes, &comment)
			items = append(items, comment)
		}
	}

	return items, nil
}
