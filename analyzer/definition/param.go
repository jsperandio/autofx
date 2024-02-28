package definition

import (
	"fmt"
	"strings"
)

// The Param struct stores information about a single parameter:
type Param struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type"`
}

// NewParam create a new Param instance from a name and type.
func NewParam(name, typ string) *Param {
	return &Param{
		Name: name,
		Type: typ,
	}
}

// String returns the parameter as a formatted string for printing
func (p Param) String() string {
	space := " "
	if p.Name == "" {
		space = ""
	}
	return fmt.Sprintf("%s%s%s", p.Name, space, p.Type)
}

// Equals returns true if the two parameters have the same name and type.
func (p *Param) Equals(other Param) bool {
	return p.Name == other.Name && p.Type == other.Type
}

// BaseType returns the base type of the parameter exclude pointer notation.
func (p *Param) BaseType() string {
	return strings.Replace(p.Type, "*", "", -1)
}
