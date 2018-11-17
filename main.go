package main

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"
)

type MyMainWindow struct {
	*walk.MainWindow
}

var inTE, outTE *walk.TextEdit
var filePath *walk.LineEdit
var mw *MyMainWindow;

func main() {

	mw = new(MyMainWindow)

	MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "SCREAMO",
		MinSize:  Size{600, 400},
		Layout:   VBox{},
		Children: []Widget{

			LineEdit{
				AssignTo: &filePath,
				Text:     SetConfig(""), //content.txt
			},
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},

			HSplitter{
				Children: []Widget{
					PushButton{
						Text:      "上传",
						OnClicked: pushGit,
					},
					PushButton{
						Text:      "打开",
						OnClicked: openMoments,
					},
				},
			},
		},
	}.Create()

	mw.Run()

}

//打开文件
func openMoments() {
	out, err := exec.Command("cmd", "/c"," start ", ""+filePath.Text()+"").CombinedOutput()
	//	out, err := exec.Command("start", "/c"," start ", ""+s+"").CombinedOutput()
	if err == nil {
		fmt.Println("start ", filePath.Text(), " err:", err)
	}
	fmt.Println("start ", filePath.Text(), " out:", string(out))
}

func pushGit() {

	var context string;

	//内容为空时，不输出
	if len(inTE.Text()) != 0 {
		context = "\n <div class=\"card hoverable\"><div class=\"card-content\"> \n"
		context += AddMarkdown(inTE.Text())
		context += "\n <p align='right'>" +
			time.Now().Format("2006-01-02") +
			"</p></div></div> \n\n"

		//输出GUI
		outTE.SetText(context)
	}

	go func() {
		//修改配置文件
		SetConfig(filePath.Text())

		//覆盖内容文本
		if len(inTE.Text()) != 0 {
			rewriteFile(filePath.Text(), context)
		}

		//提交git
		rsl := Update(filePath.Text())

		//检测提交成功否
		mw.specialAction_Triggered(rsl)
	}()
}

func (mw *MyMainWindow) specialAction_Triggered(msg string) {
	walk.MsgBox(mw, "Special", msg, walk.MsgBoxIconInformation)
}

/*
rewriteFile 覆盖写入文件
*/
func rewriteFile(file, head string) bool {
	if file == "" { //|| contents == ""
		return false
	}

	fi, err := os.Open(file) // 打开文件
	defer fi.Close()         //关闭文件
	if err != nil {
		panic(err)
	}
	contents, _ := ioutil.ReadAll(fi)                    // 读取所有内容
	newcontents := head + removePrefix(string(contents)) // 组装新的内容

	newfi, err := os.OpenFile(file, os.O_WRONLY|os.O_TRUNC, 0600)
	// newfi, err := os.OpenFile(file, os.O_RDWR, 0666) // 打开文件
	defer newfi.Close()
	if err != nil {
		panic(err)
	}
	// newfi.Seek(0, os.SEEK_SET)
	// num, err := newfi.WriteAt([]byte(newcontents), 0) // 在开头覆盖插入内容
	num, err := newfi.WriteString(newcontents) // 写入文件
	if err != nil || num < 1 {
		return false
	}
	return true
}

/*
删除旧内容首行。
1. 因为该文档要求第一行必须存在且为空。
2. 每次添加动态格式为
回车 + 动态 + 回车
动态和旧内容就是：
回车 + 动态 + 回车+回车+ 旧内容 +回车
旧内容前多了一个回车，须删除。
*/
func removePrefix(s string) string {
	tmp := strings.TrimPrefix(s, "\r\n")
	rsl := strings.TrimPrefix(tmp, "\n") //回车有区别
	return rsl
}
