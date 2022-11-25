package main

import (
	"flag"

	"github.com/kyong0612/fieldtype"

	"github.com/gqlgo/gqlanalysis/multichecker"
)

func main() {
	var fieldName string
	var typeName string
	flag.StringVar(&fieldName, "fieldName", "", "target field name")
	flag.StringVar(&typeName, "typeName", "", "expected type name")
	flag.Parse()

	multichecker.Main(
		fieldtype.Analyzer(fieldName, typeName),
	)
}
