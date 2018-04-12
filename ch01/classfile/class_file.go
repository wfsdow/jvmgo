package classfile

import "fmt"

type ClassFile struct {
	//magic uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags uint16
	thisClass uint16
	superClass  uint16
	interfaces []uint16
	fields []*MemberInfo
	methods []*MemberInfo
	attributes []AttributeInfo
}

func (self *ClassFile) read(reader *ClassReader) {
	//读取魔数
	self.readAndCheckMagic(reader)
	//读取版本号
	self.readAndCheckVersion(reader)
	//读取常量池
	self.constantPool = readConstantPool(reader)
	//读取访问标志
	self.accessFlags = reader.readUint16()
	//读取类在常量池中的索引
	self.thisClass = reader.readUint16()
	//读取超类在常量池中的索引
	self.superClass = reader.readUint16()
	//读取接口索引表
	self.interfaces = reader.readUint16s()
	//读取字段表
	self.fields = readMembers(reader, self.constantPool)
	//读取方法表
	self.fields = readMembers(reader, self.constantPool)
}



func Parse(classData []byte) (cf *ClassFile, err error){
	defer func() {
		if r :=recover(); r != nil{
			var ok bool
			err, ok = r.(error)
			if !ok{
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{}
	cf = &ClassFile{}
	cf.read(cr)
	return 

}

//判断文件的魔数是否正确
func (self *ClassFile) readAndCheckMagic(reader *ClassReader){
	magic := reader.readUint32()
	if magic != 0xCAFEBABE{
		panic("java.lang.ClassFormatError: magic!")
	}
}


//检查文件的版本
func (self *ClassFile) readAndCheckVersion(reader *ClassReader){
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45:
		return 
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion ==0{
			return
		}
	}

	panic("java.lang.UnsupportedClassVersionError!")
}

func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}
func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}
func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}
func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}
func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}
func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return ""
}

func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}
