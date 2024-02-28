package parser

import (
	"fmt"
	"go/ast"

	"github.com/jsperandio/autofx/analyzer/definition"
)

// ParseStruct function parses a Go struct from an AST type specification. It validates that the type is a struct and returns a new named struct definition.
func ParseStruct(typeSpec *ast.TypeSpec) (*definition.Struct, error) {
	_, ok := typeSpec.Type.(*ast.StructType)
	if !ok {
		return nil, fmt.Errorf("type %s is not a struct", typeSpec.Name)
	}
	return definition.NewStruct(typeSpec.Name.Name), nil
}

// ParseFunction parses a function declaration as a function. It extracts the parameters and returns a function definition.
func ParseFunction(funcDecl *ast.FuncDecl) (*definition.Function, error) {
	var err error

	functionDef := definition.NewFunction(funcDecl.Name.Name)
	functionDef.Params, err = ParseParams(funcDecl.Type.Params)
	if err != nil {
		return nil, err
	}
	functionDef.Returns, err = ParseParams(funcDecl.Type.Results)
	if err != nil {
		return nil, err
	}

	return functionDef, nil
}

// ParseInterface parses an interface type specification from the AST. It returns a new interface definition with their  methods.
func ParseInterface(typeSpec *ast.TypeSpec) (*definition.Interface, error) {
	interfaceType, ok := typeSpec.Type.(*ast.InterfaceType)
	if !ok {
		return nil, fmt.Errorf("type %s is not an interface", typeSpec.Name)
	}

	interfaceDef := definition.NewInterface(typeSpec.Name.Name)
	for _, method := range interfaceType.Methods.List {
		ft := method.Type.(*ast.FuncType)

		mtd, err := ParseMethod(ft)
		if err != nil {
			return nil, err
		}

		mtd.Name = method.Names[0].Name
		interfaceDef.Methods = append(interfaceDef.Methods, *mtd)
	}

	return interfaceDef, nil
}

// ParseMethod parses a Go method from an AST node and returns a Method definition object.
func ParseMethod(node ast.Node) (*definition.Method, error) {
	var err error

	mtd := definition.NewMethod("")
	switch n := node.(type) {
	case *ast.FuncType:

		mtd.Params, err = ParseParams(n.Params)
		if err != nil {
			return nil, err
		}

		mtd.Returns, err = ParseParams(n.Results)
		if err != nil {
			return nil, err
		}

	case *ast.FuncDecl:

		f, err := ParseFunction(n)
		if err != nil {
			return nil, err
		}

		mtd.Function = *f
		if n.Recv != nil {
			t := n.Recv.List[0].Type
			if t == nil {
				return nil, fmt.Errorf("invalid receiver type")
			}

			// mtd.SetReceiverName(n.Recv.List[0].Type.(*ast.StarExpr).X.(*ast.Ident).Name)
			mtd.SetReceiverName(getFuncDeclReceiverName(n))
		}

	default:
		return nil, fmt.Errorf("invalid node type")
	}

	return mtd, nil
}

func getFuncDeclReceiverName(funcDecl *ast.FuncDecl) string {
	if funcDecl.Recv == nil {
		return ""
	}

	if funcDecl.Recv.List[0].Type == nil {
		return ""
	}

	switch t := funcDecl.Recv.List[0].Type.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return t.X.(*ast.Ident).Name
	default:
		return ""
	}
}

// ParseParams parses parameter definitions from a Go AST FieldList and returns a slice of Param structures.
func ParseParams(fl *ast.FieldList) ([]definition.Param, error) {
	if fl == nil || len(fl.List) == 0 {
		return make([]definition.Param, 0), nil
	}

	params := make([]definition.Param, len(fl.List))
	for i := 0; i < len(fl.List); i++ {
		param := fl.List[i]
		name := ""
		if len(param.Names) > 0 {
			name = param.Names[0].Name
		}
		params[i] = *definition.NewParam(name, getPlainParamType(param.Type))
	}

	return params, nil
}

func getPlainParamType(typ ast.Expr) string {
	switch t := typ.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return "*" + getPlainParamType(t.X)
	case *ast.SelectorExpr:
		return t.X.(*ast.Ident).Name + "." + t.Sel.Name
	case *ast.ArrayType:
		return "[]" + getPlainParamType(t.Elt)
	default:
		return ""
	}
}
