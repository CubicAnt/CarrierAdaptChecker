package util

import (
	"log"
	"os/exec"
)

func RunGitCommand(workSpace string, name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	cmd.Path = workSpace
	msg, err := cmd.CombinedOutput()
	cmd.Run()
	return string(msg), err
}

func RepoSyncProject(rootDir string, project string) bool {
	_, err := RunGitCommand(rootDir, "repo", "sync", project)
	if err != nil {
		log.Println("Poject: ", project, "sync fail!!!", "error: ", err)
		return false
	}
	return true
}

func GitPullRebase(workSpace string) {
	RunGitCommand(workSpace, "git", "pull", "--rebase")
}

func GitLog(workSpace string) string {
	msg, _ := RunGitCommand(workSpace, "git", "log")
	return msg
}