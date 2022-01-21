package analyzer_test

import (
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/mbilski/exhaustivestruct/pkg/analyzer"
)

func abs(path string) string {
	s, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return s
}

func TestBasic(t *testing.T) {
	analysistest.Run(t, abs("../../testdata"), analyzer.Analyzer, "basic")
}

func TestPatternList(t *testing.T) {
	analyzer.StructPatternList = "*.Checked,*.AnotherChecked"
	analysistest.Run(t, abs("../../testdata"), analyzer.Analyzer, "patternconfig")
	analyzer.StructPatternList = ""
}

func TestInterfaceImpl(t *testing.T) {
	analyzer.StructPatternList = ""
	analysistest.Run(t, abs("../../testdata"), analyzer.Analyzer, "interfaceimpl")
}
