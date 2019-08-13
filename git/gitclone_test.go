package git

import (
	"testing"
)

func Test_gitclone(t *testing.T) {
	directory := "/home/jacob/gitdemo11"
	url := "http://msp-git.connext.com.cn:28080/PAAS-Platform/checkproperty.git"
	username := "brig"
	password := "liming12"
	branch := "testsap"
	directory = directory + "/checkproperty"

	// func gitclone(directory string, url string, username string, password string, branch string) {
	err := gitclone(directory, url, username, password, branch)
	if err != nil {
		t.Error(err.Error())
	}
}
