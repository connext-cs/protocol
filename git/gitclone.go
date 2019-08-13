package git

import (
	"errors"
	"os"
	"github.com/connext-cs/pub/logs"
	"strings"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

func gitclone(directory string, url string, username string, password string, branch string) error {
	if len(strings.TrimSpace(directory)) == 0 {
		err := errors.New("gitclone directory is ''")
		logs.Error(err)
		return err
	}

	if len(strings.TrimSpace(directory)) == 0 {
		err := errors.New("gitclone url is ''")
		logs.Error(err)
		return err
	}

	if len(strings.TrimSpace(directory)) == 0 {
		err := errors.New("gitclone username is ''")
		logs.Error(err)
		return err
	}

	if len(strings.TrimSpace(directory)) == 0 {
		err := errors.New("gitclone password is ''")
		logs.Error(err)
		return err
	}

	if len(strings.TrimSpace(branch)) == 0 {
		branch = "master"
	}

	// directory = directory + "/checkproperty"
	exit, _ := PathExists(directory)
	if exit {
		if err := os.RemoveAll(directory); err != nil {
			logs.Error(err)
		}
	}

	//git clone
	repository, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL: url,
		Auth: &http.BasicAuth{
			Username: username,
			Password: password,
		},
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		ReferenceName:     plumbing.NewBranchReferenceName(branch),
		Progress:          os.Stdout,
	})

	if err != nil {
		logs.Error(err)
		return err
	}
	logs.Info("repository:", repository)
	return nil
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
