package builders

import (
	"fmt"
	"perfe/nginxbuilder/src/types"
	"strings"
)

func Server_v1(server types.Server) string {
	serverConfig := &strings.Builder{}
	serverConfig.WriteString("server {\n") // open server code block

	if server.Secure {
		serverConfig.WriteString("\tlisten ssl;\n")
		serverConfig.WriteString(fmt.Sprintf("\tssl_certificate %s;\n", server.Ssl))
		serverConfig.WriteString(fmt.Sprintf("\tssl_certificate_key %s;\n", server.SslKey))
	} else {
		serverConfig.WriteString("\tlisten 80;\n")
	}

	serverConfig.WriteString(fmt.Sprintf("\tserver_name %s;\n", server.Name))

	serverConfig.WriteString("\tlocation ")
	if server.Location != "" {
		serverConfig.WriteString(server.Location)
	} else {
		serverConfig.WriteString("/")
	}
	serverConfig.WriteString(" {\n") // location starts here

	// Using the variable name pattern you provided earlier
	variableName := fmt.Sprintf("variable__%s__%04d", server.AsVariable(), server.ID)
	serverConfig.WriteString(fmt.Sprintf("\t\tproxy_pass http://%s;\n", variableName))
	serverConfig.WriteString("\t\tproxy_set_header Host $host;\n")
	serverConfig.WriteString("\t\tproxy_set_header X-Real-IP $remote_addr;\n")
	serverConfig.WriteString("\t\tproxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;\n")
	serverConfig.WriteString("\t\tproxy_set_header Cookie $http_cookie;\n")

	for _, prop := range server.LocationProperties {
		serverConfig.WriteString(fmt.Sprintf("\t\t%s;\n", prop))
	}

	serverConfig.WriteString("\t}\n")
	serverConfig.WriteString("}\n\n") // close server code block

	return serverConfig.String()
}
