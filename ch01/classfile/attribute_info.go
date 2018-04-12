package classfile


//jvm中属性的结构
//
//attribute_info {
//	u2 attribute_name_index;
//	u4 attribute_length;
//	u1 info[attribute_length];
//}

//属性的接口
type AttributeInfo interface {
	readInfo(read *ClassReader)
}

//读取属性表
func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributesCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributesCount)
	for i := range attributes{
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

//读取一个属性信息
func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attrNameIndex)
	attrLen := reader.readUint32()
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return  attrInfo
}

//创建具体的属性实例
func newAttributeInfo(attrName string, attrLen uint32,
	cp ConstantPool) AttributeInfo{
	switch attrName {
	case "Code": return &CodeAttribute{cp: cp}
	case "ConstantValue": return &ConstantValueAttribute{}
	case "Deprecated": return &DeprecatedAttribute{}
	case "Exceptions": return &ExceptionsAttribute{}
	case "LineNumberTable": return &LineNumberTableAttribute{}
	case "LocalVariableTable": return &LocalVariableTableAttribute{}
	case "SourceFile": return &SourceFileAttribute{cp: cp}
	case "Synthetic": return &SyntheticAttribute{}
	default: return &UnparsedAttribute{attrName, attrLen, nil}
	}
}