package fieldtype_test

import (
	"fieldtype"
	"testing"

	"github.com/gqlgo/gqlanalysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData(t)
	analysistest.Run(t, testdata, fieldtype.Analyzer("numberValue", "Numeric"), "a")
}
