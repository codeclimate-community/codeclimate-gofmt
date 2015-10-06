package main

import "github.com/codeclimate/cc-engine-go/engine"
import "sourcegraph.com/sourcegraph/go-diff/diff"
import "os/exec"
import "strconv"
import "os"
import "strings"
import "sort"

func main() {
	rootPath := "/code/"
	analysisFiles, err := engine.GoFileWalk(rootPath)
	if err != nil {
		os.Exit(1)
	}

	config, err := engine.LoadConfig()
	if err != nil {
		os.Exit(1)
	}

	excludedFiles := []string{}
	if config["exclude_paths"] != nil {
		for _, file := range config["exclude_paths"].([]interface{}) {
			excludedFiles = append(excludedFiles, file.(string))
		}
		sort.Strings(excludedFiles)
	}

	for _, path := range analysisFiles {
		cmd := exec.Command("/usr/src/app/bin/gofmt", "-d", path)
		out, err := cmd.CombinedOutput()
		if err != nil {
			return
		}

		diffs, err := diff.ParseMultiFileDiff(out)
		if err != nil {
			return
		}

		if diffs != nil && diffs[0] != nil && len(diffs[0].Hunks) > 0 {
			numHunks := strconv.Itoa((len(diffs[0].Hunks)))
			path := strings.SplitAfter(path, rootPath)[1]

			i := sort.SearchStrings(excludedFiles, path)
			if i >= len(excludedFiles) || excludedFiles[i] != path {
				issue := &engine.Issue{
					Type:              "issue",
					Check:             "GoFmt/Style/GoFmt",
					Description:       "Your code does not pass gofmt in " + numHunks + " places. Go fmt your code!",
					RemediationPoints: int32(500 * len(diffs[0].Hunks)),
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
	}

	os.Exit(0)
}
