package analyzer

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/jsperandio/autofx/analyzer/definition"
	"github.com/jsperandio/autofx/analyzer/parser"
	"github.com/jsperandio/autofx/log"
	"golang.org/x/tools/go/packages"
)

const mode packages.LoadMode = packages.NeedName |
	packages.NeedTypes |
	packages.NeedTypesInfo |
	packages.NeedSyntax

// Inspector is a struct that inspects Go packages.
type Inspector struct{}

// NewInspector returns a new Inspector instance.
func NewInspector() *Inspector {
	return &Inspector{}
}

// InspectPackage analyzes a Go package located at the given path
// and returns a Package definition.
func (i *Inspector) InspectPackage(path string) (*definition.Package, error) {
	cfg := &packages.Config{
		Fset: token.NewFileSet(),
		Mode: mode,
		Dir:  path,
	}

	loadedPackages, err := packages.Load(cfg, "")
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if len(loadedPackages) == 0 {
		log.Error("no packages found in dir %s", path)
		return nil, fmt.Errorf("no packages found")
	}

	pkg := loadedPackages[0]
	pkgdef := definition.NewPackage(pkg.Name, path)
	var mthds []*definition.Method

	for _, f := range pkg.Syntax {
		ast.Inspect(f, func(n ast.Node) bool {
			switch spec := n.(type) {
			case *ast.TypeSpec:
				log.Debug("###############################################")
				log.Debug("[Type Specification]")
				log.Debugln("")
				log.Debugf("%s ", spec.Name)

				i, err := parser.ParseInterface(spec)
				if err == nil {
					pkgdef.Interfaces[i.Name] = i
					break
				}
				log.Debugf("interface parse error: %s", err.Error())

				s, err := parser.ParseStruct(spec)
				if err == nil {
					pkgdef.Structs[s.Name] = s
					break
				}
				log.Debugf("struct parse error %s", err.Error())

			case *ast.FuncDecl:

				mthd, err := parser.ParseMethod(spec)
				if err != nil {
					log.Error("method parse error %s", err.Error())
					break
				}

				if mthd.ReceiverName() == "" {
					pkgdef.Functions[mthd.Name] = &mthd.Function
					break
				}

				if len(mthds) == 0 {
					mthds = make([]*definition.Method, 0)
				}
				mthds = append(mthds, mthd)
			}
			return true
		})
		log.Debug("###############################################")

	}

	i.methodMatch(mthds, pkgdef)
	i.constructorMatch(pkgdef)
	i.implementationsMatch(pkgdef)

	return pkgdef, nil
}

// methodMatch matches methods parsed from the AST to the respective structs .
func (*Inspector) methodMatch(mthds []*definition.Method, pkgdef *definition.Package) {
	for i := 0; i < len(mthds); i++ {
		m := mthds[i]
		s, found := pkgdef.Structs[m.ReceiverName()]
		if !found {
			continue
		}
		m.Receiver = s
		s.Methods = append(s.Methods, *m)
	}
}

// constructorMatch matches constructor functions parsed from the AST to the respective structs.
func (i Inspector) constructorMatch(pkg *definition.Package) {
	for _, f := range pkg.Functions {
		if f.IsConstructor() {
			s, found := pkg.Structs[f.Returns[0].BaseType()]
			if !found {
				continue
			}
			s.Constructor = *f
		}
	}
}

// implementationsMatch attach structs for the implemented interfaces
func (i Inspector) implementationsMatch(pkg *definition.Package) {
	for _, i := range pkg.Interfaces {
		for _, s := range pkg.Structs {
			if s.Implements(*i) {
				i.Implementations = append(i.Implementations, s.Name)
			}
		}
	}
}
