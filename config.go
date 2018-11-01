package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"fmt"
)

type Config struct {
	FileName string
}

//全局变量怎么表示
var config Config

func SetConfig(filePath string) string{
	configFile := "conf.yaml" //os.Args[1]

	source, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(source, &config)
	if err != nil {
		panic(err)
	}

	if len(filePath)!=0 {
		setAttr(filePath)
	}
	fmt.Println(config.FileName)


	out, err := yaml.Marshal(&config)
	if err != nil {
		panic(err)
	}

	err =ioutil.WriteFile(configFile,out,os.ModeAppend)
	if err != nil {
		panic(err)
	}

	return config.FileName
}

func setAttr(fileName string) {
	config.FileName=fileName
}
