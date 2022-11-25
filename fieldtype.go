package fieldtype

import (
	"fmt"

	"github.com/gqlgo/gqlanalysis"
)

func Analyzer(gqlField, gqlType string) *gqlanalysis.Analyzer {
	return &gqlanalysis.Analyzer{
		Name: "fieldtype",
		Doc:  "fieldtype finds graphql's field doesn't match type",
		Run:  run(gqlField, gqlType),
	}
}

func run(gqlField, gqlType string) func(pass *gqlanalysis.Pass) (interface{}, error) {
	return func(pass *gqlanalysis.Pass) (interface{}, error) {
		for _, t := range pass.Schema.Types {
			if t.BuiltIn {
				continue
			}
			for _, field := range t.Fields {
				if field != nil && field.Type != nil {
					fmt.Printf("📮 Read Field %s\n", field.Name)
					if field.Name == gqlField && field.Type.Name() != gqlType {
						pass.Reportf(field.Position, "%s: field %s is not %s, in fact %s",
							field.Position.Src.Name,
							gqlField,
							gqlType,
							field.Type.Name())
					}
				}
			}
		}

		return nil, nil
	}
}
