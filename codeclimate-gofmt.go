package main

import (
	"github.com/codeclimate/cc-engine-go/engine"
	"sourcegraph.com/sourcegraph/go-diff/diff"
	"os/exec"
	"strconv"
	"os"
	"strings"
	"fmt"
)

func main() {
	rootPath := "/code/"

	config, err := engine.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %s", err)
		os.Exit(1)
	}

	analysisFiles, err := engine.GoFileWalk(rootPath, engine.IncludePaths(rootPath, config))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing: %s", err)
		os.Exit(1)
	}

	for _, path := range analysisFiles {
		cmd := exec.Command("gofmt", "-d", path)

		out, err := cmd.CombinedOutput()
		if err != nil {
			return
		}

		diffs, err := diff.ParseMultiFileDiff(out)
		if err != nil {
			return
		}

		if diffs != nil && diffs[0] != nil && len(diffs[0].Hunks) > 0 {
			numHunks := len(diffs[0].Hunks)

			description := "Your code does not pass gofmt in " + strconv.Itoa(numHunks) + " " + pluralizePlace(numHunks) + ". Go fmt your code!"
			path := strings.SplitAfter(path, rootPath)[1]

			issue := &engine.Issue{
				Type:              "issue",
				Check:             "GoFmt/Style/GoFmt",
				Description:       description,
				RemediationPoints: int32(50000 * numHunks),
				Categories:        []string{"Style"},
				Location: &engine.Location{
					Path: path,
					Lines: &engine.LinesOnlyPosition{
						Begin: 1,
						End:   1,
					},
				},
			}
			engine.PrintIssue(issue)
		}
	}

	os.Exit(0)
}

func pluralizePlace(quant int) string {
	if quant > 1 {
		return "places"
	} else {
		return "place"
	}
}
