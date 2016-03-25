package main

import (
	"fmt"

	"github.com/voxelbrain/goptions"
)

const (
	PUBKEY_NAME  = "%v"
	PRIVKEY_NAME = "%v.priv"
)

func main() {
	options := struct {
		User string        `goptions:"-u, --username, obligatory, description='Username to authenticate with'"`
		Key string         `goptions:"-i, --identity, obligatory, description='SSH Key to authenticate with'"`
		Project string     `goptions:"-p, --project, obligatory, description='Project to deploy against'"`
		Runner string     `goptions:"-r, --runner, obligatory, description='Runner to which to deploy'"`

		goptions.Verbs
		Deploy struct {
		} `goptions:"deploy"`
		Clean struct {
		} `goptions:"clean"`
		Build struct {
		} `goptions:"build"`
		Run struct {
		} `goptions:"run"`
		Stdin struct {
			File string `goptions:"-f, --file, description='File of desired STDIN'"`
		} `goptions:"stdin"`
		Stdout struct {
		} `goptions:"stdout"`
	}{
	}

	fmt.Println("CodeRunnerConsole-NG v1")
	goptions.ParseAndFail(&options)

	if len(options.Verbs) <= 0 {
		fmt.Println("You must specify a verb")
		return
	}

	switch options.Verbs {
	case "deploy":
		doDeploy(options.User, options.Key, options.Project, options.Runner)
		break
	case "clean":
		doClean(options.User, options.Key, options.Project, options.Runner)
		break
	case "build":
		doBuild(options.User, options.Key, options.Project, options.Runner)
		break
	case "run":
		doRun(options.User, options.Key, options.Project, options.Runner)
		break
	case "stdin":
		doStdin(options.User, options.Key, options.Project, options.Runner, options.Stdin.File)
		break
	case "stdout":
		doStdout(options.User, options.Key, options.Project, options.Runner)
		break
	}
}
