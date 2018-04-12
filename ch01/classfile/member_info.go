package classfile

type MemberInfo struct {
	cp ConstantPool
	accessFlags uint16
	nameIndex uint16
	descriptorInfo uint16
	attributes []AttributeInfo
}
