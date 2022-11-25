package main

import (
	"flag"
	"fmt"

	"github.com/kyong0612/fieldtype"

	"github.com/gqlgo/gqlanalysis/multichecker"
)

func main() {
	fmt.Println("##### START FIELD #####")

	// fieldtype -fieldName="numverValue" -typeName="Numeric" -schema="sidecar/gql/schema/**/*.graphql"

	// go run cmd/fieldtype/main.go -fieldName="numverValue" -typeName="Numeric" -schema="server/graphql/schema/**/*.graphql"

	var fieldName string
	var typeName string
	flag.StringVar(&fieldName, "fieldName", "", "target field name")
	flag.StringVar(&typeName, "typeName", "", "expected type name")
	flag.Parse()

	multichecker.Main(
		fieldtype.Analyzer("numberValue", "Numeric"),
	)
}
