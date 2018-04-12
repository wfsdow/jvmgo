package classfile

//表示废弃的属性
type DeprecatedAttribute struct {
	MarkerAttribute
}

type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct {}

func (slef *MarkerAttribute) readInfo(reader *ClassReader){

}