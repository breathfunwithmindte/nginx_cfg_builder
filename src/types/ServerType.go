package types

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

type Server struct {
	ID                 int      `json:"-"`
	Kind               string   `json:"kind"`
	Name               string   `json:"name"`
	Secure             bool     `json:"secure"`
	Builder            string   `json:"builder"`
	Ssl                string   `json:"ssl"`          // example: /etc/letsencrypt/live/ecommerce-template-v1.perfect-evolution.com/cert.pem
	SslKey             string   `json:"ssl_key"`      //  /etc/letsencrypt/live/ecommerce-template-v1.perfect-evolution.com/privkey.pem
	LocationProperties []string `json:"l_properties"` // list<proxy_set_header Cookie $http_cookie>
	Skip               bool     `json:"skip"`
	Location           string   `json:"location"`
	LocalServers       []string `json:"local_servers"`
	Description        string   `json:"description"`
}

func (s *Server) Init() {
	rand.Seed(time.Now().UnixNano())
	s.ID = rand.Intn(10000) // Generate a random 4-digit number
}

func (s Server) Type() string {
	return "Server"
}

func (s Server) AsVariable() string {
	req := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	variable := req.ReplaceAllString(s.Name, "_")
	variable = strings.ToLower(variable)
	return variable
}

func (s Server) PrettyPrint() {
	blue := "\033[34m"
	reset := "\033[0m"
	red := "\033[31m"
	green := "\033[32m"

	fmt.Printf(
		"\tID: %s %d %s | Type: %s %s %s | Builder: %s %s %s | Name %s %s %s \n",
		blue, s.ID, reset,
		red, s.Type(), reset,
		blue, s.Builder, reset,
		green, fmt.Sprintf(" %s | %s ", s.Name, s.AsVariable()), reset)

	fmt.Println("---------------------------------------------------------------------")
}

func (s Server) Build() string {
	builderFunc := GetServerBuilder(s.Builder)
	if s.Skip == true {
		return ""
	}
	if builderFunc != nil {
		return builderFunc(s)
	}
	return fmt.Sprintf("# [SERVER] => Builder not found with name := %s \n", s.Builder) // Return an comment line explaining the error
}
