package main

import (
	"encoding/csv"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

var dir = ""

func gitCommand(command string) (result string) {
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
		result += fmt.Sprintf("Git Error : %s \n ", err)
	} else {
		result += fmt.Sprintf("Git Result : %s \n ", out)
	}
	return
}

func Update(filePath string) (rsl string){

	path, name := filepath.Split(filePath)
	dir = path

	//fmt.Printf("dir=%s , name = %s \n",dir,name)

	rsl += gitCommand("add " + name)
	rsl += gitCommand("commit -m \"update moments 更新动态\"")
	rsl += gitCommand("push origin master")

	return
}
