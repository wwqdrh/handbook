package container

type Person struct {
	Age int
}

func (p *Person) GrowUp() {
	p.Age++
}

func MapUpdate() {
	// m := map[string]Person{
	// 	"iswbm": Person{Age: 20},
	// }
	// m["iswbm"].Age = 23 // will panic
	// m["iswbm"].GrowUp() // will panic

	m := map[string]*Person{
		"iswbm": {Age: 20},
	}
	m["iswbm"].GrowUp()
}
