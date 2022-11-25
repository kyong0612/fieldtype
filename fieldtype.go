package fieldtype

import (
	"fmt"

	"github.com/gqlgo/gqlanalysis"
)

func Analyzer(gqlField, gqlType string) *gqlanalysis.Analyzer {
	return &gqlanalysis.Analyzer{
		Name: "fieldtype",
		Doc:  "fieldtype finds field to not have input type",
		Run:  run(gqlField, gqlType),
	}
}

func run(gqlField, gqlType string) func(pass *gqlanalysis.Pass) (interface{}, error) {
	return func(pass *gqlanalysis.Pass) (interface{}, error) {

		// fmt.Printf("👿 ast.Definition %#v\n", pass.Schema.Mutation.Position)

		for _, t := range pass.Schema.Types {
			if t.BuiltIn {
				continue
			}
			for _, field := range t.Fields {
				if field != nil && field.Type != nil {
					fmt.Printf("📮 field %#v\n", *field)
					if field.Name == gqlField && field.Type.Name() != gqlType {
						pass.Reportf(field.Position, "field %s is not %s", gqlField, gqlType)
						pass.Reportf(field.Position, "field %s is %s", field.Name, field.Type.Name())
						return nil, fmt.Errorf("%s: field %s is %s", field.Position.Src.Name, field.Name, field.Type.Name())
						// return nil, fmt.Errorf("field %s is %s", field.Name, field.Type.Name())
					}
					// if !field.Type.NonNull {
					// 	if field.Position != nil {
					// 		pass.Reportf(field.Position, "%s is optional", field.Name)
					// 	}
					// }
					// for _, argument := range field.Arguments {
					// 	if argument != nil && argument.Type != nil {
					// 		if !argument.Type.NonNull {
					// 			pass.Reportf(argument.Position, "%s is optional", argument.Name)
					// 		}
					// 	}
					// }
				}
			}
		}

		return nil, nil
	}
}

// func checkOperations(pass *gqlanalysis.Pass, ops ast.OperationList) {
// 	for _, op := range ops {
// 		for _, sel := range op.SelectionSet {
// 			checkField(pass, sel)
// 			checkInlineFragment(pass, sel)
// 		}
// 	}
// }

// func checkFragments(pass *gqlanalysis.Pass, fs ast.FragmentDefinitionList) {
// 	for _, f := range fs {
// 		checkFragment(pass, f)
// 	}
// }

// func checkFragment(pass *gqlanalysis.Pass, f *ast.FragmentDefinition) {
// 	if !hasID(f.Definition.Fields) {
// 		return
// 	}

// 	var hasID bool
// 	for _, s := range f.SelectionSet {
// 		hasID = hasID || isID(s)
// 		checkField(pass, s)
// 	}

// 	if hasID {
// 		return
// 	}

// 	name := f.Definition.Name
// 	if name == "" {
// 		name = "fragment"
// 	}

// 	pass.Reportf(f.Position, "type %s has id field but %s does not have id field", name, f.Name)
// }

// func checkField(pass *gqlanalysis.Pass, sel ast.Selection) {
// 	field, _ := sel.(*ast.Field)
// 	if field == nil || allFragmentSpread(field.SelectionSet) {
// 		return
// 	}

// 	for _, s := range field.SelectionSet {
// 		checkField(pass, s)
// 		checkInlineFragment(pass, s)
// 	}

// 	ft := pass.Schema.Types[field.Definition.Type.Name()]
// 	if ft == nil || !hasID(ft.Fields) {
// 		return
// 	}

// 	var hasID bool
// 	for _, s := range field.SelectionSet {
// 		hasID = hasID || isID(s)
// 	}

// 	if !hasID && !allInlineFragment(field.SelectionSet) {
// 		pass.Reportf(field.Position, "type %s has id field but selection %s does not have id field", ft.Name, field.Name)
// 	}
// }

// func hasID(fields ast.FieldList) bool {
// 	for _, field := range fields {
// 		if field.Name == "id" {
// 			return true
// 		}
// 	}
// 	return false
// }

// func isID(sel ast.Selection) bool {
// 	field, _ := sel.(*ast.Field)
// 	return field != nil && field.Name == "id"
// }

// func checkInlineFragment(pass *gqlanalysis.Pass, sel ast.Selection) {
// 	f, _ := sel.(*ast.InlineFragment)
// 	if f == nil || allFragmentSpread(f.SelectionSet) {
// 		return
// 	}

// 	for _, s := range f.SelectionSet {
// 		checkField(pass, s)
// 		checkInlineFragment(pass, s)
// 	}

// 	if !hasID(f.ObjectDefinition.Fields) {
// 		return
// 	}

// 	var hasID bool
// 	for _, s := range f.SelectionSet {
// 		hasID = hasID || isID(s)
// 	}

// 	if !hasID && !allInlineFragment(f.SelectionSet) {
// 		pass.Reportf(f.Position, "type %s has id field but fragment does not have id field", f.ObjectDefinition.Name)
// 	}
// }

// func allFragmentSpread(set ast.SelectionSet) bool {
// 	for _, sel := range set {
// 		if _, ok := sel.(*ast.FragmentSpread); !ok {
// 			return false
// 		}
// 	}
// 	return true
// }

// func allInlineFragment(set ast.SelectionSet) bool {
// 	for _, sel := range set {
// 		if _, ok := sel.(*ast.InlineFragment); !ok {
// 			return false
// 		}
// 	}
// 	return true
// }
