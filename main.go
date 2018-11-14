package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main() {
	var inTE, outTE *walk.TextEdit
	var filePath *walk.LineEdit

	mainWindow := MainWindow{
		Title:   "SCREAMO",
		MinSize: Size{600, 400},
		Layout:  VBox{},
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

			PushButton{
				Text: "SCREAM",
				OnClicked: func() {
					var context = "\n <div class=\"card hoverable\"><div class=\"card-content\"> \n"
					context += AddMarkdown(inTE.Text())
					context += "\n <p align='right'>" +
						time.Now().Format("2006-01-02") +
						"</p></div></div> \n\n"

					//输出GUI
					outTE.SetText(context)

					go func() {
						//输出文件
						SetConfig(filePath.Text())
						rewriteFile(filePath.Text(), context)

						//提交git
						Update(filePath.Text())
					}()
				},
			},
		},
	}

	mainWindow.Run()

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
	tmp:=strings.TrimPrefix(s, "\r\n")
	rsl:=strings.TrimPrefix(tmp, "\n")//回车有区别
	return rsl
}
