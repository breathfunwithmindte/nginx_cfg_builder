package types

// Element represents any element in the configuration.
type Element interface {
	Type() string
	PrettyPrint()
	Build() string
}
