package main

import (
	"regexp"
	"log"
)

func doValidateUsername(username string) bool {
	r, err := regexp.Compile("[^a-zA-Z0-9]")

	if err != nil {
		log.Println(err)
		return false
	}

	n_username := r.ReplaceAllString(username, "")

	if n_username != username {
		return false
	}

	return true
}

func doValidateKey(key string) bool {
	// TODO - validate path exists
	return true
}

func doValidateProject(project string) bool {
	// Merp
	r, err := regexp.Compile("[^a-zA-Z0-9_/-]")

	if err != nil {
		log.Println(err)
		return false
	}

	n_project := r.ReplaceAllString(project, "")

	if n_project != n_project {
		return false
	}

	return true
}

func doValidateRunner(runner string) bool {
	if (runner != "runner1" && runner != "runner2") {
		return false
	}
	return true
}

func doValidateAll(username, key, project, runner string) bool {
	if (!doValidateUsername(username)) {
		log.Println("Invalid username.")
		return false
	}
	if (!doValidateKey(key)) {
		log.Println("Invalid key.")
		return false
	}
	if (!doValidateProject(project)) {
		log.Println("Invalid project.")
		return false
	}
	if (!doValidateRunner(runner)) {
		log.Println("Invalid runner. Choices: runner1, runner2")
		return false
	}
	return true
}
