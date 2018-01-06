package aux

import (
	"os"
	
	"os/exec"
	"bytes"
)

func Tar(source, target string) error {
	args := []string{"-cvzf", target, source}
	cmd := exec.Command("tar", args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func Untar(tarball, targetDir string) error {
	
	os.Chdir(targetDir)
	args := []string{"-xvzf", tarball}
	
	cmd := exec.Command("tar", args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}