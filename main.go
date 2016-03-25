package main

import (
	"fmt"
	"log"

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
		Runner string      `goptions:"-r, --runner, obligatory, description='Runner to which to deploy'"`
		Help goptions.Help `goptions:"-h, --help, description='Show this help'"`

		goptions.Verbs
		Deploy struct {
		} `goptions:"deploy"`
		Clean struct {
		} `goptions:"clean"`
		Build struct {
		} `goptions:"build"`
		Run struct {
			Redirect string `goptions:"-x, --redirect, description='Which redirect: normal, none, all'"`
		} `goptions:"run"`
		Stdin struct {
			File string `goptions:"-f, --file, obligatory, description='File to set STDIN. If not specified, prints STDIN'"`
		} `goptions:"stdin"`
		Get struct {
			Method string `goptions:"-m, --method, obligatory, description='Method to fetch: stdin, stdout, or stderr'"`
		} `goptions:"get"`
	}{
		Redirect: "normal",
	}

	log.Println("crconsole-ng v1")
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
		doRun(options.User, options.Key, options.Project, options.Runner, options.Run.Redirect)
		break
	case "stdin":
		doStdin(options.User, options.Key, options.Project, options.Runner, options.Stdin.File)
		break
	case "get":
		doGet(options.User, options.Key, options.Project, options.Runner, options.Get.Method)
		break
	}
}
