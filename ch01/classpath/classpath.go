package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry //启动类入口
	extClasspath  Entry //拓展类入口
	userClasspath Entry //自定义类入口
}

//解析三种类的路径，
//-Xjre选项解析启动类、拓展类路径
//-classpath/-cp选项解析用户类路径
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

//依次从启动类路径，拓展类路径、用户类路径中搜索class文件
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}

	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}

	return self.userClasspath.readClass(className)
}

func (self *Classpath) String() string {
	return self.userClasspath.String()
}

//解析启动类，拓展类路径，将内容保存到Classpath对应的字段中
func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	//生成 jre/lib/* 的路径
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)
	//生成 jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

//解析用户类路径
func (self *Classpath) parseUserClasspath(cpOption string) {
	//如果没有提供-classpath/-cp选项，则将当前目录作为用户类路径
	if cpOption == "" {
		cpOption = "."
	}

	self.userClasspath = newEntry(cpOption)
}

//获取jre目录
func getJreDir(jreOption string) string {
	//获取-Xjre选项中的路径
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}

	//获取当前目录中的jre路径
	if exists("./jre") {
		return "./jre"
	}

	//获取环境变量JAVA_HOME
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}

	panic("Can not find jre floder")

}

//判断目录是否存在
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
