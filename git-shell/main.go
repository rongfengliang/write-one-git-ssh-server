package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func init() {
	file := "./" + "message" + ".txt"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
}
func parseSSHCmd(cmd string) (string, string) {
	ss := strings.SplitN(cmd, " ", 2)
	if len(ss) != 2 {
		return "", ""
	}
	return ss[0], strings.Replace(ss[1], "'/", "'", 1)
}
func main() {
	println(os.Getenv("SSH_ORIGINAL_COMMAND"))
	sshCmd := os.Getenv("SSH_ORIGINAL_COMMAND")
	verb, args := parseSSHCmd(sshCmd)
	repoFullName := strings.ToLower(strings.Trim(args, "'"))
	verbs := strings.Split(verb, " ")
	var gitCmd *exec.Cmd
	gitRepoPath := fmt.Sprintf("/opt/gitrepo/%s", repoFullName)
	os.Setenv("MY_UID", "dalongrong")
	if len(verb) == 2 {
		println(verbs[0], verbs[1], gitRepoPath)
		gitCmd = exec.Command(verbs[0], verbs[1], gitRepoPath)
	} else {
		println(verbs[0], gitRepoPath)
		gitCmd = exec.Command(verb, gitRepoPath)
	}
	gitCmd.Stdin = os.Stdin
	gitCmd.Stdout = os.Stdout
	gitCmd.Stderr = os.Stderr
	if err := gitCmd.Run(); err != nil {
		log.Fatal("Internal error", "Failed to execute git command: %v", err)
	}
	return
}
