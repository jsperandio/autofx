package main

import (
	"flag"

	"github.com/jsperandio/autofx/analyzer"
	"github.com/jsperandio/autofx/generator"
	"github.com/jsperandio/autofx/log"
)

var (
	pkgPathFlag  *string
	logLevelFlag *string
)

func flagParse() {
	pkgPathFlag = flag.String("p", "", "package path")
	logLevelFlag = flag.String("ll", "info", "log level")
	flag.Parse()

	if *pkgPathFlag == "" {
		log.Error("flag [-p] is required")
		panic("Package path is required")
	}

	if *logLevelFlag != "info" && *logLevelFlag != "debug" {
		log.Error("flag [-ll] is invalid")
		panic("Invalid log level")
	}
}

func main() {
	flagParse()
	log.Init(logLevelFlag)

	ins := analyzer.NewInspector()
	def, err := ins.InspectPackage(*pkgPathFlag)
	if err != nil {
		log.Fatal(err)
	}
	// def.Print()

	gen := generator.NewGenerator(def)
	err = gen.Generate()
	if err != nil {
		log.Error(err)
	}
}
