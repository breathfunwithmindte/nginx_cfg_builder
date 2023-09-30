package builders

import (
	"perfe/nginxbuilder/src/types"
)

func Comment_v1(c types.Comment) string {
	// For the sake of this example, let's assume the `Comment` type has a Docs field which is a slice of strings.
	// Convert this slice into a single string where each comment line is prefixed with a '# '.
	commentStr := ""
	for _, line := range c.Docs {
		commentStr += "# " + line + "\n"
	}
	return commentStr + "\n"
}
