package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func openInVSCode(folder string) bool {
	if isDirExist(folder) {
		log.Println("Git Repo already exist. Opening in VS-Code...")
		exec.Command("code", folder).Output()
		return true
	}
	return false
}

func isDirExist(dirName string) bool {
	info, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		log.Println("Dir not exits ", dirName)
		return false
	}
	return info.IsDir()
}

func removeDir(dirName string) {
	os.RemoveAll(dirName)
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

	if ok := openInVSCode(lastPath); ok {
		os.Exit(0)
	}

	log.Println("Git Remote Repo availablity checking...")

	_, err := exec.Command("git", "ls-remote", path).Output()
	if err != nil {
		log.Fatal("Not a valid git repo: ", path)
	}

	log.Println("Git Remote Repo Downloading...")

	_, err = exec.Command("git", "clone", path).Output()
	if err != nil {
		log.Fatal("Not able to do git clone from ", path)
	}

	log.Println("Git Repo Downloaded. Opening in VS-Code...")
	openInVSCode(lastPath)

	return true
}
