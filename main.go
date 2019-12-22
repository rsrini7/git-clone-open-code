package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func openInVSCode(folder string) {
	_, err := exec.Command("code", folder).Output()
	if err != nil {
		log.Fatal("Not able to open vscode from ", folder, err)
	}
}

func isDirExist(dirName string) bool {
	info, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		log.Println("Dir not exits ", dirName)
		return false
	}
	return info.IsDir()
}

func main() {
	parseArgsAndExecute()
}

func parseArgsAndExecute() bool {
	var path string
	if len(os.Args) != 2 {
		log.Fatal("Should have git url as second arg")
		//path = "https://github.com/rsrini7/MyElmResume"
	}

	path = os.Args[1]

	lastPath := path[strings.LastIndex(path, "/")+1:]
	if isDirExist(lastPath) {
		log.Println("Git Repo already exist. Opening in VS-Code...")
		openInVSCode(lastPath)
		os.Exit(0)
	} else {
		log.Println("Git Repo Downloading...")
	}

	_, err := exec.Command("git", "ls-remote", path).Output()
	if err != nil {
		log.Fatal("Not a valid git repo: ", path)
	}

	_, err = exec.Command("git", "clone", path).Output()
	if err != nil {
		log.Fatal("Not able to do git clone from ", path)
	}

	if isDirExist(lastPath) {
		log.Println("Git Repo Downloaded. Opening in VS-Code...")
		openInVSCode(lastPath)
	}
	return true
}
