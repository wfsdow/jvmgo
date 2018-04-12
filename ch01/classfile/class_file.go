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
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
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