package definition

import (
	"fmt"
	"strings"
	"unicode"
)

// Function struct defines a Go function with its name, whether it is private, its parameters and return values.
type Function struct {
	Name    string  `json:"name"`
	Private bool    `json:"-"`
	Params  []Param `json:"params,omitempty"`
	Returns []Param `json:"returns,omitempty"`
}

// NewFunction method returns a new Function object.
func NewFunction(name string) *Function {
	return &Function{
		Name:    name,
		Private: len(name) > 0 && unicode.IsLower(rune(name[0])),
	}
}

// Signature method returns a string representation of the function signature including its name,
// parameters and return values formatted as expected in Go code.
func (f *Function) Signature() string {
	returns := stringfyParam(f.Returns)
	if len(f.Returns) > 1 {
		returns = fmt.Sprintf("(%s)", returns)
	}

	return fmt.Sprintf("%s(%s) %s", f.Name, stringfyParam(f.Params), returns)
}

// IsConstructor check if the function is a constructor(starts with New).
func (f *Function) IsConstructor() bool {
	return strings.HasPrefix(f.Name, "New")
}

// stringfyParam method returns a string representation of a slice of Param objects.
// Ex:
//
//	[]Param{{"a", "int"}, {"b", "string"}} -> "a int, b string"
func stringfyParam(slice []Param) string {
	var params []string
	for _, p := range slice {
		params = append(params, p.String())
	}
	return strings.Join(params, ",")
}
