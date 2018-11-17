package main

import (
	"fmt"
	"os/exec"
	"testing"
)

func Test_gitCommand(t *testing.T) {
	//exec.Command("cmd", "/c", " ping ", "127.0.0.1")
	//cmd := exec.Command("cmd","/c","ping 127.0.0.1")
	cmd := exec.Command("cmd","/c","git push origin master")
	cmd.Dir = "F:\\GitHub\\caliburn1994.github.io\\_includes\\"

	var rsl string
	out, err := cmd.CombinedOutput()
	if err != nil {
		rsl += fmt.Sprintf("Git Error : %s \n ", err)
	} else {
		rsl += fmt.Sprintf("Git Result : %s \n ", out)
	}

	fmt.Println(rsl)
}
