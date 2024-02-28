package definition

// Struct struct stores information about a Go struct definition
type Struct struct {
	Name        string   `json:"name"`
	Methods     []Method `json:"methods,omitempty"`
	Constructor Function `json:"constructor,omitempty"`
}

// NewStruct create a new Struct instance from a name. Initializes empty slices for Methods
func NewStruct(name string) *Struct {
	return &Struct{
		Name:    name,
		Methods: make([]Method, 0),
	}
}

// Type return the type name of struct
func (s Struct) Type() string {
	return s.Name
}

// Implements Checks if a struct implements an interface by comparing method names and signatures.
func (s *Struct) Implements(iface Interface) bool {
	for _, mtd := range iface.Methods {
		strcMtd := s.getMethodByName(mtd.Name)
		if strcMtd == nil {
			return false
		}

		if !mtd.Same(*strcMtd) {
			return false
		}
	}
	return true
}

// getMethodByName searches the Methods slice for a method with a matching name.
func (s *Struct) getMethodByName(name string) *Method {
	for _, method := range s.Methods {
		if method.Name == name {
			return &method
		}
	}
	return nil
}
