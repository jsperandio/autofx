package definition

import (
	"fmt"
	"slices"
)

// Method is a struct that represents a method in a Go struct. It contains the name of the method, the parameters, and the return values. It also contains a pointer to the Struct that the method belongs to.
type Method struct {
	Function
	Receiver     *Struct `json:"-"`
	receiverName string  `json:"-"`
}

// NewMethod function initializes a new Method struct with the given name.
func NewMethod(name string) *Method {
	return &Method{
		Function: Function{
			Name:    name,
			Params:  make([]Param, 0),
			Returns: make([]Param, 0),
		},
	}
}

// ReceiverName returns the name of the receiver of the method.
func (m *Method) ReceiverName() string {
	return m.receiverName
}

// SetReceiverName sets the name of the receiver of the method.
func (m *Method) SetReceiverName(name string) {
	m.receiverName = name
}

// Same compares two Method structs and returns true if their definition matches. This would be useful for identifying duplicate methods.
func (m *Method) Same(check Method) bool {
	if m.Name != check.Name {
		return false
	}
	if len(m.Params) != len(check.Params) {
		return false
	}
	if len(m.Returns) != len(check.Returns) {
		return false
	}

	for i := range check.Params {
		if !slices.Contains(m.Params, check.Params[i]) {
			return false
		}
	}

	for i := range check.Returns {
		if !slices.Contains(m.Returns, check.Returns[i]) {
			return false
		}
	}

	return true
}

// Signature returns the signature of the method. This is the same as the signature of the Function struct. The receiver name is prepended to the signature. This is useful for generating documentation.
func (m *Method) Signature() string {
	receiverName := m.ReceiverName()
	if m.Receiver != nil {
		receiverName = m.Receiver.Name
	}
	return fmt.Sprintf("(%s) %s", receiverName, m.Function.Signature())
}
