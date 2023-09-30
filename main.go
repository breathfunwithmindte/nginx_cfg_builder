package main

import (
	"log"
	"perfe/nginxbuilder/src/services"
)

func main() {
	configsDirectory := "configs/" // Path to your configs directory

	configs, err := services.LoadConfigsFromDir(configsDirectory)
	if err != nil {
		log.Fatalf("Error loading configurations: %v", err)
	}

	services.RunProcess(configs)

	// for _, config := range configs {
	// 	config.PrettyPrint()
	// }
}
