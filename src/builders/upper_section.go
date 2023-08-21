package builders

import (
	"fmt"
	t "perfe/nginxbuilder/src/types"
)

func GenerateUpperSection(servers []t.Server) string {
	upperSection := "# Auto-generated Nginx configuration by perfe/nginxbuilder\n"
	upperSection += "# Do not modify directly, as changes might be overwritten.\n\n"
	upperSection += "# Server upstream variables:\n"

	for _, server := range servers {
		variableName := fmt.Sprintf("variable__%s__%04d", server.AsVariable(), server.ID)
		upperSection += fmt.Sprintf("upstream %s {", variableName)
		for _, localServer := range server.LocalServers {
			upperSection += fmt.Sprintf(" server %s;", localServer)
		}
		upperSection += " }\n"
	}

	return upperSection + "\n"
}
