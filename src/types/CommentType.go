package types

import (
	"fmt"
)

type Comment struct {
	Kind    string   `json:"kind"`
	Builder string   `json:"builder"`
	Docs    []string `json:"docs"`
}

func (c Comment) Type() string {
	return "Comment"
}

func (c Comment) PrettyPrint() {
	blue := "\033[34m"
	reset := "\033[0m"
	red := "\033[31m"
	yellow := "\033[33m"
	fmt.Printf("\tType: %s %s %s | Builder: %s %s %s \n", red, c.Type(), reset, blue, c.Builder, reset)

	for _, doc := range c.Docs {
		fmt.Printf("\t\t\t @ %s %s %s\n", yellow, doc, reset)
	}
	fmt.Println("---------------------------------------------------------------------")
}

func (c Comment) Build() string {
	builderFunc := GetCommentBuilder(c.Builder)
	if builderFunc != nil {
		return builderFunc(c)
	}
	return fmt.Sprintf("# [Comment] => Builder not found with name := %s \n", c.Builder) // Return an comment line explaining the error
}
