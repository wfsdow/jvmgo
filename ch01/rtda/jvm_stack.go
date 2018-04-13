package rtda

//使用链表实现的栈
type Stack struct {
	maxSize uint  //栈的容量
	size uint  //栈当前的大小
	_top *Frame  //栈顶指针
}

//创建一个栈，指定能容纳的最大帧数量
func newStack(maxSize uint) *Stack{
	return &Stack{
		maxSize:  maxSize,
	}
}

//向栈中压入一个帧
func (self *Stack) push(frame *Frame) {
	//判断栈是否满了
	if self.size >= self.maxSize{
		panic("java.lang.StackOverflowError")
	}

	if self._top !=nil{
		frame.lower = self._top
	}
	self._top = frame
	self.size ++
}

//从栈中弹出一个帧
func (self *Stack) pop(frame *Frame) *Frame{
	//判断是否为空栈
	if  self._top == nil{
		panic("jvm stack is empty!")
	}

	top := self._top
	self._top = top.lower
	top.lower = nil
	self.size --
	return top
}
