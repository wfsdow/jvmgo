package classfile

//无法解析的属性
type UnparsedAttribute struct {
	name string
	length uint32
	info []byte
}

func (self *UnparsedAttribute) readInfo(reader *ClassReader){
	self.info = reader.readBytes(self.length)
}