package classfile

//存放方法的行号信息
//LineNumberTable_attribute {
//	u2 attribute_name_index;
//	u4 attribute_length;
//	u2 line_number_table_length;
//	{ 	u2 start_pc;
//		u2 line_number;
//	} line_number_table[line_number_table_length];
//}

type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc uint16
	lineNumber uint16
}

func (self *LineNumberTableAttribute) readInfo(reader *ClassReader){
	lineNumberTableLength := reader.readUint16()
	self.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range self.lineNumberTable{
		self.lineNumberTable[i] = &LineNumberTableEntry{
			startPc: reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}