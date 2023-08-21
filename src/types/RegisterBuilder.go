package types

// Define the function signature for our builders
type CommentBuilderFunc func(Comment) string
type ServerBuilderFunc func(Server) string

// Our centralized registry
var commentBuilders = make(map[string]CommentBuilderFunc)
var serverBuilders = make(map[string]ServerBuilderFunc)

// Setters
func RegisterCommentBuilder(version string, builder CommentBuilderFunc) {
	commentBuilders[version] = builder
}
func RegisterServerBuilder(version string, builder ServerBuilderFunc) {
	serverBuilders[version] = builder
}

// Getters
func GetCommentBuilder(version string) CommentBuilderFunc {
	return commentBuilders[version]
}
func GetServerBuilder(version string) ServerBuilderFunc {
	return serverBuilders[version]
}
