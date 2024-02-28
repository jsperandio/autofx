package template

type ModuleData struct {
	PackageName          string
	ConstructorName      string
	ImplementPackageName string
	ImplementType        string
}

const (
	GoFileInits = `package {{.}}

import "go.uber.org/fx"
`

	InterfaceModule = `
func {{.ImplementType}}Module() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotate(
				{{ if .PackageName }}{{.PackageName}}.{{end}}{{.ConstructorName}},
				fx.As(new({{ if .ImplementPackageName }}{{.ImplementPackageName}}.{{end}}{{.ImplementType}})),
			),
		),
	)
}
`

	SimpleModule = `
func {{.ImplementType}}Module() fx.Option {
	return fx.Options(
		fx.Provide(
			{{.ConstructorName}},
		),
	)
}
`

	PackageModule = `
func Module() fx.Option {
	return fx.Options(
	{{range .}}	{{ if .ImplementPackageName }}{{.ImplementPackageName}}.{{end}}{{.ImplementType}}Module(),
	{{end}})
}
`
)
