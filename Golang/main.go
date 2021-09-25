package main

type People interface {
	Say()
}

func DoSay(p People) {
	p.Say()
}

type parent struct {
	Name string
}

func (p *parent) Say() {
	println(p.Name)
}

type Child struct {
	parent
}

func (p *Child) Say() {
	println(p.Name)
}

func main() {
	people := &Child{parent: parent{Name: "zs"}}
	println(people.Name)
	people.Name = "ls"
	println(people.Name)
	println(people.parent.Name)

	a := &parent{Name: "parent"}
	b := &Child{parent: parent{Name: "child"}}
	a.Say()
	b.Say()

	DoSay(a)
	DoSay(b)
}
