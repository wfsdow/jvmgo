package main

import "fmt"
import "strings"
import "jvmgo/ch01/classpath"

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	//解析三种类路径
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath: %v class: %v args:%v\n",
		cp, cmd.class, cmd.args)
	//将类名中的.替换为/
	className := strings.Replace(cmd.class, ".", "/", -1)
	//读取类文件
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
		return
	}
	fmt.Printf("class data:%v\n", classData)
}
