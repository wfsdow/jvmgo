package rtda

type Thread struct {
	pc int
	stack *Stack
}

func NewThread() *Thread{
	return &Thread{
		stack: newStack(1024)
	}
}

func (self *Thread) PushFrame(frame * Frame) {
	self.stack.push(frame)
}

func (self *Thread) PopFrame(frame * Frame) {
	self.stack.pop(frame)
}
