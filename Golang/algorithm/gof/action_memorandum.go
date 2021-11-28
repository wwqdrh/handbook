package gof

import "fmt"

// Originator 发起者，负责创建备忘录，可以记录和恢复本身的状态。
type Originator struct {
	// State 当前状态
	State string
}

// Memento 备忘录，存储发起者的状态，防止除发起者之外的人访问备忘录
type Memento struct {
	// state 当前状态，不可导出
	state string
}

// Caretaker 管理者，负责存储备忘录，但不对备忘录操作
type Caretaker struct {
	// memento 备忘录对象
	memento Memento
}

func NewCaretaker(mem Memento) *Caretaker {
	return &Caretaker{
		memento: mem,
	}
}

// restoreMemento 将发起人的状态更新为备忘录的状态，即恢复至备忘录
func (originator *Originator) RestoreMemento(creator *Caretaker) {
	originator.State = creator.memento.state
}

// createMemento 使用发起者当前的状态，创建备忘录
func (originator *Originator) CreateMemento() Memento {
	return Memento{originator.State}
}

// print 打印发起者的状态
func (originator *Originator) Print() {
	fmt.Println(fmt.Sprintf("originator的状态为%s", originator.State))
}
