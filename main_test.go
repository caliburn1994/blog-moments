package main

import (
	"fmt"
	"os/exec"
	"testing"
)

//TODO 测试，包内存在依赖
func Test_openMoments(t *testing.T) {
	s := "F:\\GitHub\\caliburn1994.github.io\\_includes\\moments.md"

	//cmd := exec.Command("cmd", "/c", "start \"\" \"c:\\some path\\withspaces\\init\\thefile.pptx\"")
	/*env := os.Environ()
	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	pid, err :=os.StartProcess(s,[]string{},procAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!", err)  //
		os.Exit(1)
	}
	fmt.Printf("The process id is %v", pid)*/
	//	out, err := exec.Command("start", "/c"," start ", ""+s+"").CombinedOutput()
	out, err := exec.Command("cmd", "/c", " ping ", "127.0.0.1").CombinedOutput()

	if err == nil {
		fmt.Println("start ", s, " err:", err)
	}
	fmt.Println("start ", s, " out:", string(out))

}
