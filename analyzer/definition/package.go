package definition

import (
	"encoding/json"
	"fmt"
)

// Package struct defines a Go package with its name and maps containing any Interfaces, Structs and Functions it contains.
// The maps are indexed by their name.
//
// Example:
//
//	package main
//
//	import (
//		"github.com/ryan-berger/go-def"
//	)
//
//	func main() {
//		p := def.NewPackage("main")
//		p.Interfaces["Interface"] = def.NewInterface("Interface")
//		p.Structs["Struct"] = def.NewStruct("Struct")
//		p.Functions["Function"] = def.NewFunction("Function")
//		p.Report()
//	}
type Package struct {
	Name       string                `json:"name"`
	Path       string                `json:"path"`
	Interfaces map[string]*Interface `json:"interfaces,omitempty"`
	Structs    map[string]*Struct    `json:"structs,omitempty"`
	Functions  map[string]*Function  `json:"functions,omitempty"`
}

// NewPackage function initializes a new Package struct with the given name. It initializes the type maps to empty maps to allow types to be added later.
func NewPackage(name string, path string) *Package {
	return &Package{
		Name:       name,
		Path:       path,
		Interfaces: make(map[string]*Interface, 0),
		Structs:    make(map[string]*Struct, 0),
		Functions:  make(map[string]*Function, 0),
	}
}

// Print method prints a formatted report of the contents of the package including statistics on the number of each type.
func (p *Package) Print() {
	var (
		clReset   = "\033[0m"
		clrGreen  = "\033[32m"
		clrYellow = "\033[33m"
		clrRed    = "\033[31m"
		clrBlue   = "\033[34m"
		clrPurple = "\033[35m"
	)

	fmt.Printf("==================================================\n\n")
	fmt.Printf("  Package  %s %s%s\n\n", clrGreen, p.Name, clReset)
	fmt.Printf("%s  %sInterfaces(%d)%s----------------------------%s\n", clrYellow, clReset, len(p.Interfaces), clrYellow, clReset)
	for _, i := range p.Interfaces {
		fmt.Printf("%s   %s%s\n", clrGreen, i.Name, clReset)
		for _, f := range i.Methods {
			fmt.Printf("    %s%s%s%s\n", clrYellow, "├", f.Name, clReset)
		}
		fmt.Printf("    %s%s%s\n", clrPurple, "Implementations", clReset)
		for _, s := range p.Structs {
			if s.Implements(*i) {
				fmt.Printf("     %s%s%s%s\n", clrGreen, "└", s.Name, clReset)
			}
		}
	}
	fmt.Printf("%s  -----------------------------------------%s\n", clrYellow, clReset)
	fmt.Printf("\n%s  %sStructs(%d)%s-------------------------------%s\n", clrRed, clReset, len(p.Structs), clrRed, clReset)
	for _, s := range p.Structs {
		fmt.Printf("%s   %s%s\n", clrGreen, s.Name, clReset)
		fmt.Printf("%s    %s%s%s\n", clrBlue, "╚", s.Constructor.Name, clReset)
		for _, m := range s.Methods {
			fmt.Printf("    %s%s%s%s\n", clrYellow, "├", m.Name, clReset)
			// fmt.Printf("%s   %s%s\n", clrYellow, m.Signature(), clReset)
		}
	}

	fmt.Printf("%s  -----------------------------------------%s\n", clrRed, clReset)
	fmt.Printf("\n%s  %sFunctions(%d)%s-----------------------------%s\n", clrPurple, clReset, len(p.Functions), clrPurple, clReset)
	for _, f := range p.Functions {
		fmt.Printf("%s   %s%s\n", clrYellow, f.Name, clReset)
		// fmt.Printf("%s   %s%s\n", clrYellow, f.Signature(), clReset)
	}
	fmt.Printf("%s  -----------------------------------------%s\n", clrPurple, clReset)

	fmt.Printf("==================================================\n")
}

// AsJSON method returns a JSON representation of the Package struct.
func (p *Package) AsJSON() string {
	b, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return ""
	}
	return string(b)
}
