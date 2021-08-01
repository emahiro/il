package main

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestMain(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, a, "a")
}
