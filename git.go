package main

import (
	"encoding/csv"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

var dir string =""

func gitCommand(command string) {
	// Split string 分割参数
	r := csv.NewReader(strings.NewReader(command))
	r.Comma = ' ' // space
	paras, err := r.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	//run
	cmd := exec.Command("git", paras...)
	cmd.Dir = dir


	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("git error : %s ",err)
	}

	//转换windows cmd的编码
	fmt.Println("Result:", string(out))
}



func Update(filePath string) {

	path ,name:=filepath.Split(filePath)
	dir=path

	fmt.Printf("dir=%s , name = %s \n",dir,name)

	gitCommand("add "+name)
	gitCommand("commit -m \"update moments 更新动态\"")
	gitCommand("push origin master")
	//	PushGit()
}

/*func main() {
	Update("F:/GitHub/caliburn1994.github.io/_includes/moments.md")
}*/


