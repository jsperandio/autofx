package definition

// Interface struct defines a Go interface with its name and a slice of Methods.
type Interface struct {
	Name            string   `json:"name"`
	Methods         []Method `json:"methods"`
	Implementations []string `json:"implementations"`
}

// NewInterface function initializes a new Interface struct with the given name. It sets the Methods field to an empty slice to allow methods to be added later.
func NewInterface(name string) *Interface {
	return &Interface{
		Name:            name,
		Methods:         make([]Method, 0),
		Implementations: make([]string, 0),
	}
}

// Type function returns the name of the interface.
func (i Interface) Type() string {
	return i.Name
}
