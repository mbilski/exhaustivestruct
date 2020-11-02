package analyzer_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/mbilski/exhaustivestruct/pkg/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAll(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}

	testdata := filepath.Join(filepath.Dir(filepath.Dir(wd)), "testdata")
	analyzer.StructPatternList = "*.Test,*.Test2,*.Embedded,*.External"
	analysistest.Run(t, testdata, analyzer.Analyzer, "s")
}
