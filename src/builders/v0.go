package builders

import (
	"fmt"
	t "perfe/nginxbuilder/src/types"
	"strings"
)

func Builder_v1(server *t.Server, variable_name string) *strings.Builder {

	serverConfig := &strings.Builder{}
	serverConfig.WriteString("server {\n") // open server code block
	if server.Secure {
		serverConfig.WriteString("\tlisten ssl;\n") //\tlisten 127.0.0.1:443 ssl;\n
		serverConfig.WriteString(fmt.Sprintf("\tssl_certificate %s;\n", server.Ssl))
		serverConfig.WriteString(fmt.Sprintf("\tssl_certificate_key %s;\n", server.SslKey))
	} else {
		serverConfig.WriteString("\tlisten 80;\n")
	}

	serverConfig.WriteString(fmt.Sprintf("\tserver_name %s;\n", server.Name))
	// if server.ServerLine != "default" {
	// 	serverConfig.WriteString(fmt.Sprintf("\t%s;\n\n", server.ServerLine))
	// } else {
	// 	serverConfig.WriteString("\tlimit_req zone=one burst=3 nodelay;\n\n")
	// }

	serverConfig.WriteString("\tlocation / {\n") // location starts here
	// whatever goes inside location / { ... }
	serverConfig.WriteString(fmt.Sprintf("\t\tproxy_pass http://%s;\n", variable_name))
	serverConfig.WriteString("\t\tproxy_set_header Host $host;\n")
	serverConfig.WriteString("\t\tproxy_set_header X-Real-IP $remote_addr;\n")
	serverConfig.WriteString("\t\tproxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;\n")
	serverConfig.WriteString("\t\tproxy_set_header Cookie $http_cookie;\n")
	serverConfig.WriteString("\t}\n\n")
	serverConfig.WriteString("}\n") // close server code block

	return serverConfig

}
