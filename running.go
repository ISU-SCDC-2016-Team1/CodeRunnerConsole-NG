package main

import (
	"log"
	"os/exec"
	"crypto/md5"
	"io"
	"fmt"
	"bytes"
)

func doReturnRunnerIP(runner string) string {
	if (runner == "runner1") {
		return "10.2.2.2"
	} else if (runner == "runner2") {
		return "10.2.2.3"
	}

	return ""
}

func doForwardIdentity(username, key, project, runner string) {
	if (!doValidateAll(username, key, project, runner)) {
		return
	}

	log.Printf("Forward Identity %v to %v as %v with %v...\n", project, runner, username, key)

	cmd := exec.Command("scp", "-i", key, key, fmt.Sprintf("%v@%v:~/.ssh/id_rsa", username, doReturnRunnerIP(runner)))
	cmd.Run()
}

func doRemoveForwardedIdentity(username, key, project, runner string) {
	if (!doValidateAll(username, key, project, runner)) {
		return
	}

	log.Printf("Remove Identity %v to %v as %v with %v...\n", project, runner, username, key)

	cmd := exec.Command("ssh", "-i", key, key, fmt.Sprintf("%v@%v", username, doReturnRunnerIP(runner)), "rm", "-rf", "~/.ssh/id_rsa")
	cmd.Run()
}

func doDeploy(username, key, project, runner string) {
	if (!doValidateAll(username, key, project, runner)) {
		return
	}

	doForwardIdentity(username, key, project, runner)

	log.Printf("Deploying %v to %v as %v with %v...\n", project, runner, username, key)

	h := md5.New()
	io.WriteString(h, project)
	var md5sum string = string(h.Sum(nil))

	cmd := exec.Command("ssh", "-i", key, fmt.Sprintf("%v@%v", username, doReturnRunnerIP(runner)), "mkdir", "-p", "~/deploy/" + md5sum, " && ",  "cd", "~/deploy/" + md5sum, " && ", "git clone", "https://git.team1.isucdc.com/" + project)
	cmd.Run()

	doRemoveForwardedIdentity(username, key, project, runner)
}

func doClean(username, key, project, runner string) {
	if (!doValidateAll(username, key, project, runner)) {
		return
	}

	log.Printf("Cleaning %v to %v as %v with %v...\n", project, runner, username, key)

	h := md5.New()
	io.WriteString(h, project)
	var md5sum string = string(h.Sum(nil))

	cmd := exec.Command("ssh", "-i", key, fmt.Sprintf("%v@%v", username, doReturnRunnerIP(runner)), "rm", "-rf", "~/deploy/" + md5sum)
	cmd.Run()

	doRemoveForwardedIdentity(username, key, project, runner)
}

func doBuild(username, key, project, runner string) {
	if (!doValidateAll(username, key, project, runner)) {
		return
	}
	log.Printf("Building %v to %v as %v with %v...\n", project, runner, username, key)

	h := md5.New()
	io.WriteString(h, project)
	var md5sum string = string(h.Sum(nil))

	var out bytes.Buffer
	cmd := exec.Command("ssh", "-i", key, fmt.Sprintf("%v@%v", username, doReturnRunnerIP(runner)), "cd", "~/deploy/" + md5sum + "/*", " && ", "bash ./.build")
	cmd.Stdout = &out
	cmd.Run()
	fmt.Println(out.String())
}

func doRun(username, key, project, runner, method string) {
	if (!doValidateAll(username, key, project, runner)) {
		return
	}

	if (!doValidateMethod(method)) {
		return
	}
	log.Printf("Run %v to %v as %v with %v...\n", project, runner, username, key)

	h := md5.New()
	io.WriteString(h, project)
	var md5sum string = string(h.Sum(nil))

	var out bytes.Buffer
	if (method == "normal") {
		cmd := exec.Command("ssh", "-i", key, fmt.Sprintf("%v@%v", username, doReturnRunnerIP(runner)), "cd ~/deploy/" + md5sum + "/* && bash ./.run > .stdout 2>/dev/null")
		cmd.Stdout = &out
		cmd.Run()
	} else if (method == "all") {
		cmd := exec.Command("ssh", "-i", key, fmt.Sprintf("%v@%v", username, doReturnRunnerIP(runner)), "cd ~/deploy/" + md5sum + "/* && bash ./.run >.stdout 2>.stderr <.stdin")
		cmd.Stdout = &out
		cmd.Run()
	} else {
		cmd := exec.Command("ssh", "-i", key, fmt.Sprintf("%v@%v", username, doReturnRunnerIP(runner)), "cd ~/deploy/" + md5sum + "/* && bash ./.run >/dev/null 2>/dev/null")
		cmd.Stdout = &out
		cmd.Run()
	}

	fmt.Println(out.String())
}

func doStdin(username, key, project, runner, stdin string) {
	if (!doValidateAll(username, key, project, runner)) {
		return
	}

	if (!doValidateFile(stdin)) {
		log.Printf("File does not exist:", stdin)
		return
	}

	log.Printf("Stdin %v to %v as %v with %v...\n", project, runner, username, key)

	h := md5.New()
	io.WriteString(h, project)
	var md5sum string = string(h.Sum(nil))

	cmd := exec.Command("scp", "-i", key, stdin, fmt.Sprintf("%v@%v:~/.stdin", username, doReturnRunnerIP(runner)))
	cmd.Run()

	cmd = exec.Command("ssh", "-i", key, fmt.Sprintf("%v@%v", username, doReturnRunnerIP(runner)), "cd ~/deploy/" + md5sum + "/* && mv ~/.stdin ./.stdin")
	cmd.Run()
}

func doGet(username, key, project, runner, method string) {
	if (!doValidateAll(username, key, project, runner)) {
		return
	}

	if (!doValidateMethod(method)) {
		return
	}

	log.Printf("Fetching %v : %v to %v as %v with %v...\n", method, project, runner, username, key)

	h := md5.New()
	io.WriteString(h, project)
	var md5sum string = string(h.Sum(nil))

	var out bytes.Buffer
	cmd := exec.Command("ssh", "-i", key, fmt.Sprintf("%v@%v", username, doReturnRunnerIP(runner)), "cd ~/deploy/" + md5sum + "/* && cat ./." + method)
	cmd.Stdout = &out
	cmd.Run()

	fmt.Println(out.String())
}
