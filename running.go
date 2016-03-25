package main


func doDeploy(username, key, project, runner string) {
	if (!doValidateAll(username, key, project, runner)) {
		return
	}
}

func doClean(username, key, project, runner string) {
	if (!doValidateAll(username, key, project, runner)) {
		return
	}
}

func doBuild(username, key, project, runner string) {
	if (!doValidateAll(username, key, project, runner)) {
		return
	}
}

func doRun(username, key, project, runner string) {
	if (!doValidateAll(username, key, project, runner)) {
		return
	}
}

func doStdin(username, key, project, runner, stdin string) {
	if (!doValidateAll(username, key, project, runner)) {
		return
	}
}

func doStdout(username, key, project, runner string) {
	if (!doValidateAll(username, key, project, runner)) {
		return
	}
}
