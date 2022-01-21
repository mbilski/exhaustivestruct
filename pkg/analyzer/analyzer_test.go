package analyzer_test

import (
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/mbilski/exhaustivestruct/pkg/analyzer"
)

func testdata() string {
	s, err := filepath.Abs("../../testdata")
	if err != nil {
		panic(err)
	}
	return s
}

func TestBasic(t *testing.T) {
	analysistest.Run(t, testdata(), analyzer.Analyzer, "basic")
}

func TestPatternList(t *testing.T) {
	analyzer.StructPatternList = "*.Checked,*.AnotherChecked"
	analysistest.Run(t, testdata(), analyzer.Analyzer, "patternconfig")
	analyzer.StructPatternList = ""
}

func TestInterfaceImpl(t *testing.T) {
	analyzer.StructPatternList = ""
	analysistest.Run(t, testdata(), analyzer.Analyzer, "interfaceimpl")
}
