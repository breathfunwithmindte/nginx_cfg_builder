package services

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"perfe/nginxbuilder/src/builders"
	"perfe/nginxbuilder/src/types"
	"strings"
)

// RunProcess is the primary function that processes the list of Config structures.
func RunProcess(configs []types.Config) error {
	// Registering builders
	types.RegisterCommentBuilder("v1", builders.Comment_v1)
	types.RegisterServerBuilder("v1", builders.Server_v1)

	for _, config := range configs {
		var configFile strings.Builder

		var servers []types.Server
		for _, item := range config.Items {

			if server, ok := item.(types.Server); ok {
				servers = append(servers, server)
			}
		}

		// Add the upper section
		configFile.WriteString(builders.GenerateUpperSection(servers))

		// Loop through the items in each Config
		for _, item := range config.Items {
			configFile.WriteString(item.Build())
			item.PrettyPrint()
			// TODO: Incorporate logic to generate middle section based on item type
		}

		// Add the bottom section
		configFile.WriteString(builders.GenerateBottomSection())

		// Save the configFile content to an actual file in the /build directory
		outputPath := filepath.Join("build", config.Output+".conf")
		if err := ioutil.WriteFile(outputPath, []byte(configFile.String()), os.ModePerm); err != nil {
			return fmt.Errorf("failed to save config to %s: %v", outputPath, err)
		}
	}
	return nil
}
