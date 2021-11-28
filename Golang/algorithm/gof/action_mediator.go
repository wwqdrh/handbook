package gof

import "fmt"

type Colleague interface {
	Send(msg string)
	Notify(msg string)
	SetMediator(mediator Mediator)
}

type Mediator interface {
	Send(msg string, colleague Colleague)
}

type ConcreteColleague1 struct {
	mediator Mediator
}
type ConcreteColleague2 struct {
	mediator Mediator
}

func (c *ConcreteColleague1) SetMediator(mediator Mediator) {
	c.mediator = mediator
}

func (c *ConcreteColleague1) Send(msg string) {
	c.mediator.Send(msg, c)
}

func (c *ConcreteColleague1) Notify(msg string) {
	fmt.Println("ConcreteColleague1 recv msg:", msg)
}

func (c *ConcreteColleague2) SetMediator(mediator Mediator) {
	c.mediator = mediator
}

func (c *ConcreteColleague2) Send(msg string) {
	c.mediator.Send(msg, c)
}

func (c *ConcreteColleague2) Notify(msg string) {
	fmt.Println("ConcreteColleague2 recv msg:", msg)
}

type ConcreteMediator struct {
	C1 Colleague
	C2 Colleague
}

func (c *ConcreteMediator) Send(msg string, colleague Colleague) {
	if colleague == c.C1 {
		c.C2.Notify(msg)
	} else {
		c.C1.Notify(msg)
	}
}
