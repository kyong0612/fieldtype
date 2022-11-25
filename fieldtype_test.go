package fieldtype_test

import (
	"testing"

	"github.com/kyong0612/fieldtype"

	"github.com/gqlgo/gqlanalysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData(t)
	analysistest.Run(t, testdata, fieldtype.Analyzer("numberValue", "Numeric"), "a")
}
