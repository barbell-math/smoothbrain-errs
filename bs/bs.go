package main

import (
	"bytes"
	"context"
	"strings"

	sbbs "github.com/barbell-math/smoothbrain-bs"
)

func gitDiffStage(errMessage string, targetToRun string) sbbs.StageFunc {
	return sbbs.Stage(
		"Run Diff",
		func(ctxt context.Context, cmdLineArgs ...string) error {
			var buf bytes.Buffer
			if err := sbbs.Run(ctxt, &buf, "git", "diff"); err != nil {
				return err
			}
			if buf.Len() > 0 {
				sbbs.LogErr(errMessage)
				sbbs.LogQuietInfo(buf.String())
				sbbs.LogErr(
					"Run build system with %s and push any changes",
					targetToRun,
				)
				return sbbs.StopErr
			}
			return nil
		},
	)
}

func main() {
	sbbs.RegisterTarget(
		context.Background(),
		"updateDeps",
		sbbs.Stage(
			"barbell math package cmds",
			func(ctxt context.Context, cmdLineArgs ...string) error {
				var packages bytes.Buffer
				if err := sbbs.Run(
					ctxt, &packages, "go", "list", "-m", "-u", "all",
				); err != nil {
					return err
				}

				lines := strings.Split(packages.String(), "\n")
				// First line is the current package, skip it
				for i := 1; i < len(lines); i++ {
					iterPackage := strings.SplitN(lines[i], " ", 2)
					if !strings.Contains(iterPackage[0], "barbell-math") {
						continue
					}

					if err := sbbs.RunStdout(
						ctxt, "go", "get", iterPackage[0]+"@latest",
					); err != nil {
						return err
					}
				}
				return nil
			},
		),
		sbbs.Stage(
			"Non barbell math package cmds",
			func(ctxt context.Context, cmdLineArgs ...string) error {
				if err := sbbs.RunStdout(ctxt, "go", "get", "-u", "./..."); err != nil {
					return err
				}
				if err := sbbs.RunStdout(ctxt, "go", "mod", "tidy"); err != nil {
					return err
				}

				return nil
			},
		),
	)
	sbbs.RegisterTarget(
		context.Background(),
		"updateReadme",
		sbbs.Stage(
			"Run gomarkdoc",
			func(ctxt context.Context, cmdLineArgs ...string) error {
				err := sbbs.RunStdout(
					ctxt, "gomarkdoc", "--output", "README.md", ".",
				)
				if err != nil {
					sbbs.LogQuietInfo("Consider running build system with installGoMarkDoc target if gomarkdoc is not installed")
				}
				return err
			},
		),
	)
	sbbs.RegisterTarget(
		context.Background(),
		"installGoMarkDoc",
		sbbs.Stage(
			"Install gomarkdoc",
			func(ctxt context.Context, cmdLineArgs ...string) error {
				return sbbs.RunStdout(
					ctxt, "go",
					"install", "github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest",
				)
			},
		),
	)
	sbbs.RegisterTarget(
		context.Background(),
		"fmt",
		sbbs.Stage(
			"Run go fmt",
			func(ctxt context.Context, cmdLineArgs ...string) error {
				return sbbs.RunStdout(ctxt, "go", "fmt", "./...")
			},
		),
	)
	sbbs.RegisterTarget(
		context.Background(),
		"unitTests",
		sbbs.Stage(
			"Run go test",
			func(ctxt context.Context, cmdLineArgs ...string) error {
				return sbbs.RunStdout(ctxt, "go", "test", "-v", "./...")
			},
		),
	)

	sbbs.RegisterTarget(
		context.Background(),
		"ciCheckDeps",
		sbbs.TargetAsStage("updateDeps"),
		gitDiffStage("Out of date packages were detected", "updateDeps"),
	)
	sbbs.RegisterTarget(
		context.Background(),
		"ciCheckReadme",
		sbbs.TargetAsStage("installGoMarkDoc"),
		sbbs.TargetAsStage("updateReadme"),
		gitDiffStage("Readme is out of date", "updateReadme"),
	)
	sbbs.RegisterTarget(
		context.Background(),
		"ciCheckFmt",
		sbbs.TargetAsStage("fmt"),
		gitDiffStage("Fix formatting to get a passing run!", "fmt"),
	)
	sbbs.RegisterTarget(
		context.Background(),
		"mergegate",
		sbbs.TargetAsStage("ciCheckFmt"),
		sbbs.TargetAsStage("ciCheckReadme"),
		sbbs.TargetAsStage("ciCheckDeps"),
		sbbs.TargetAsStage("unitTests"),
	)

	sbbs.Main("build")
}
