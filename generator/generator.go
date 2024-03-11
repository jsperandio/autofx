package generator

import (
	"fmt"
	"text/template"

	"github.com/jsperandio/autofx/analyzer/definition"
	tmpl "github.com/jsperandio/autofx/generator/template"
)

const defaultFileName = "module.go"

type Generator struct {
	Package *definition.Package

	dependencies  map[*definition.Struct][]string
	resultModules []string
	resultFile    *File
}

func NewGenerator(pkg *definition.Package) *Generator {
	if pkg == nil {
		return nil
	}

	return &Generator{
		Package:       pkg,
		dependencies:  make(map[*definition.Struct][]string),
		resultFile:    NewFile(defaultFileName, pkg.Path, nil),
		resultModules: []string{},
	}
}

func (g *Generator) Generate() error {
	g.buildDependencyMap()

	err := g.initFileContent()
	if err != nil {
		return err
	}

	err = g.fillSimpleTemplates()
	if err != nil {
		return err
	}

	err = g.fillInterfaceTemplates()
	if err != nil {
		return err
	}

	err = g.fillPackageModule()
	if err != nil {
		return err
	}

	_, err = g.resultFile.Save()
	if err != nil {
		return err
	}

	return nil
}

func (g *Generator) buildDependencyMap() {
	for _, v := range g.Package.Structs {
		if len(v.Constructor.Params) == 0 {
			g.dependencies[v] = nil
			continue
		}
		deps := make([]string, len(v.Constructor.Params))
		for i, p := range v.Constructor.Params {
			deps[i] = p.BaseType()
		}
		g.dependencies[v] = deps
	}
}

func (g *Generator) initFileContent() error {
	t := template.Must(template.New("init").Parse(tmpl.GoFileInits))
	err := t.Execute(g.resultFile, g.Package.Name)
	if err != nil {
		return err
	}

	return nil
}

func (g *Generator) fillSimpleTemplates() error {
	t := template.Must(template.New("simpleModule").Parse(tmpl.SimpleModule))

	for dep, needs := range g.dependencies {
		if needs != nil || len(needs) > 0 {
			continue
		}

		md := tmpl.ModuleData{
			PackageName:     g.Package.Name,
			ConstructorName: dep.Constructor.Name,
			ImplementType:   dep.Name,
		}

		err := t.Execute(g.resultFile, md)
		if err != nil {
			return err
		}

		g.resultModules = append(g.resultModules, dep.Name)
	}

	return nil
}

func (g *Generator) fillInterfaceTemplates() error {
	t := template.Must(template.New("interfaceModule").Parse(tmpl.InterfaceModule))

	for _, ifc := range g.Package.Interfaces {
		md := tmpl.ModuleData{
			ImplementType: ifc.Type(),
		}

		for _, impl := range ifc.Implementations {
			s, ok := g.Package.Structs[impl]
			if !ok {
				return fmt.Errorf("struct not found %s", impl)
			}
			md.ConstructorName = s.Constructor.Name
		}

		err := t.Execute(g.resultFile, md)
		if err != nil {
			return err
		}

		g.resultModules = append(g.resultModules, ifc.Type())
	}

	return nil
}

func (g *Generator) fillPackageModule() error {
	t := template.Must(template.New("packageModule").Parse(tmpl.PackageModule))

	mdL := make([]tmpl.ModuleData, len(g.resultModules))
	for i, m := range g.resultModules {
		mdL[i] = tmpl.ModuleData{
			ImplementType: m,
		}
	}

	err := t.Execute(g.resultFile, mdL)
	if err != nil {
		return err
	}

	return nil
}
